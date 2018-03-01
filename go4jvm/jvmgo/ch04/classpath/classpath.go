package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

/*
1. 类加载的顺序
1.1 bootstrap 加载的jre/lib/*路径
1.2 ext 加载的jre/lib/ext/*路径
1.3 用户指定的类路径
*/
func Parse(jreOption string, cpOption string) *Classpath {
	classPath := &Classpath{}
	classPath.parseBootAndExtClasspath(jreOption)
	classPath.parseUserClasspath(cpOption)
	return classPath
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 先获取jre目录
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	// Xjre参数
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	// 当前目录
	if exists("./jre") {
		return "./jre"
	}

	// JAVA_HOME
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return jh
	}

	panic("Can not find jre folder")
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	// 如果cpOption没有设置，则默认当前目录
	if cpOption == "" {
		cpOption = "."
	}
	// 因为cpOption是外界传入，所以无法确定classpath的类型
	self.userClasspath = newEntry(cpOption)
}

// 判断路径是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		// 如果获取路径对应的信息错误
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 在使用具体类时，指定的是包名+类名；在具体读取这个class文件时，需要补全文件扩展名
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
