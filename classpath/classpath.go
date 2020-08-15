package classpath

import (
	"os"
	"path/filepath"
)

//Classpath 对应类的路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//Parse 解析文件目录
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//ReadClass 读取类
func (client *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := client.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := client.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return client.userClasspath.readClass(className)
}

func (client *Classpath) String() string {
	return client.userClasspath.String()
}

func (client *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	client.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	client.extClasspath = newWildcardEntry(jreExtPath)
}

func (client *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	client.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
