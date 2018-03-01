package classpath

import (
	"path/filepath"
	"io/ioutil"
)

/*
Go语言中结构体不需要显示的实现接口，
只要结构体的方法匹配则代表实现了接口;
需要方法全部匹配成功，才代表实现；
*/
/*用于表示目录类路径*/
type DirEntry struct {
	//用于存放绝对路径
	absDir string
}

/*
Go语言没有专门的构造函数;
统一使用new开头的函数模拟构造器，生成结构体实例
*/
func newDirEntry(path string) *DirEntry {
	// 将path转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 终止程序的执行
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([] byte, Entry, error) {
	//将相对路径变成绝对路径
	fileName := filepath.Join(self.absDir, className)
	data, error := ioutil.ReadFile(fileName)
	return data, self, error
}

func (self *DirEntry) String() string {
	return self.absDir
}
