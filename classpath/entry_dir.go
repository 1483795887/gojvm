package classpath

import "path/filepath"
import "io/ioutil"

//DirEntry 表示目录类型的entry
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (client *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(client.absDir,className)
	data, err := ioutil.ReadFile(fileName)
	return data, client, err
}

func (client *DirEntry) String() string                                    {
	return client.absDir
}
