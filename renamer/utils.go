package renamer

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GetFileExt 获取文件扩展名
func GetFileExt(filePath string) (ext string) {
	if fileInfo, err := os.Stat(filePath); err == nil && !fileInfo.IsDir() {
		ext = strings.ToLower(path.Ext(filePath))
	}
	return
}

// GetFileName 获取文件名
func GetFileName(filePath string) (name string) {
	if fileInfo, err := os.Stat(filePath); err == nil && !fileInfo.IsDir() {
		name = fileInfo.Name()
		ext := path.Ext(filePath)
		name = strings.TrimRight(name, ext)
	}
	return
}

// FilePathWalkDir 遍历目录
func FilePathWalkDir(root string) (files map[string]os.FileInfo, dirPath []string, err error) {
	files = make(map[string]os.FileInfo)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path != root {
			if !info.IsDir() {
				files[path] = info
			} else {
				dirPath = append(dirPath, path)
			}
		}
		return nil
	})
	return
}
