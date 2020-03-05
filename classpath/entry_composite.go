package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 就是 Entry列表
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	// 按照分隔符分开所有的的类路径，然后用newEntry加载他们
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (self CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {

	for _, entry := range self {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (self CompositeEntry) String() string {

	strs := make([]string, len(self))

	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
