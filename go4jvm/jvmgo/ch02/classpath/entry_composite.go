package classpath

import "strings"

type CompositeEntry []Entry

func newCompsiteEntry(pathList string) CompositeEntry {
	// 定义一个slice
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	return nil, nil, nil
}

func (self CompositeEntry) String() string {
	return ""
}
