package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ZipEntry 表示zip或jar格式的entry
type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

func (client *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(client.absDir)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, client, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (client *ZipEntry) String() string {
	return client.absDir
}
