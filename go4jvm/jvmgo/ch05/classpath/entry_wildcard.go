package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		/*
		对通配符类路径，只接受处理路径下的.jar结尾的jar文件；
		对子目录不进行处理
		*/
		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry2(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
