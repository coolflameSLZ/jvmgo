package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption string, cpOption string) *ClassPath {

	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := self.bootClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}

	return self.userClassPath.ReadClass(className)
}

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}

func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	// 第一优先级,用户输入的目录存在
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 第二优先级,当前目录下有名叫jre的文件
	if exists("./jre") {
		return ".jre"
	}
	// 第三优先级，找找环境变量中有没有java路径
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("can not find jre folder!没有找到jre目录")
}

//exists（）函数用于判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 将用户传入的classpath路径传入 结构体
func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.userClassPath = newEntry(cpOption)
}
