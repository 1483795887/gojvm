package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

//Entry 是表示不同层级的文件
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path,".zip")||strings.HasSuffix(path,".ZIP"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
