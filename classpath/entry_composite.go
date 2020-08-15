package classpath

import (
	"errors"
	"strings"
)

//CompositeEntry 表示组合entry
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (client CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range client {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (client CompositeEntry) String() string {
	strs := make([]string, len(client))
	for i, entry := range client {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
