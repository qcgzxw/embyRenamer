package renamer

import (
	"os"
	"path"
	"path/filepath"
	"strconv"
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
		name = name[:len(name)-len(filepath.Ext(name))]
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

// GetNumStr 获取剧集编号
func GetNumStr(total uint, no string) (numStr string) {
	numStr = no
	if len(strconv.Itoa(int(total))) <= 1 {
		total = 10
	}
	if len(strconv.Itoa(int(total))) > len(no) {
		numStr = strings.Repeat("0", len(strconv.Itoa(int(total)))-len(no)) + no
	}
	return
}

func OsRename(oldPath, newPath string) error {
	if strings.ToLower(GetFileExt(oldPath)) == ".nfo" {
		if _, ok := invalidNfoPath[oldPath]; ok {
			delete(invalidNfoPath, oldPath)
		}
	}
	return os.Rename(oldPath, newPath)
}
