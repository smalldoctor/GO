package classpath

import "path/filepath"

type ZipEntry struct {
	// 用来存放zip或者jar文件的绝对路径
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([] byte, Entry, error) {
	return nil, nil, nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}
