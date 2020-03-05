package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // 去掉路径最后的 *
	compositeEntry := []Entry{}
	// 将路径中的所有jar包都装置进 compositeEntry
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 过滤非baseDir的文件，因为*不能递归扫描
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
