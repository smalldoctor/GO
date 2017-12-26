package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/*
	className是class的全限定名(因为包路径+类名就是class的相对路径名)，
	相对路径，文件以.class结尾;以 / 分割;
	Entry是指向具体的class;
	*/
	readClass(className string) ([] byte, Entry, error)
	// 相当于Java中的toString方法
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompsiteEntry(path)
	}
}
