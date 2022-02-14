package renamer

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
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

// FilePathWalkDir 深度遍历目录
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

// FilePathWalkCurrentDir 遍历当前目录
func FilePathWalkCurrentDir(root string) (files map[string]os.DirEntry, dirPath []string, err error) {
	files = make(map[string]os.DirEntry)
	dirFiles, err := os.ReadDir(root)
	if err != nil {
		return
	}
	for _, file := range dirFiles {
		p := root + string(os.PathSeparator) + file.Name()
		if file.IsDir() {
			dirPath = append(dirPath, p)
		} else {
			files[p] = file
		}
	}
	return
}

// CountFileInDir 获取当前目录下某后缀文件的数量
func CountFileInDir(dirPath, fileExt string) (num int) {
	if files, _, err := FilePathWalkDir(dirPath); err == nil && len(files) > 0 {
		for filePath, _ := range files {
			if GetFileExt(filePath) == fileExt {
				num++
			}
		}
	}
	return
}

// GetNumStr 获取剧集编号
func GetNumStr(total uint, no string) (numStr string) {
	numStr = no
	if len(strconv.Itoa(int(total))) <= 1 {
		total = 10
	}
	if no == "0" {
		return strings.Repeat("0", len(strconv.Itoa(int(total))))
	}
	if len(strconv.Itoa(int(total))) > len(no) {
		numStr = strings.Repeat("0", len(strconv.Itoa(int(total)))-len(strings.TrimLeft(no, "0"))) + no
	}
	return
}

// OsRename 重命名
func OsRename(oldPath, newPath, newPathRoot string) error {
	if strings.ToLower(GetFileExt(oldPath)) == ".nfo" {
		if _, ok := invalidNfoPath[oldPath]; ok {
			delete(invalidNfoPath, oldPath)
		}
	}
	if _, err := os.Stat(oldPath); err != nil {
		// 旧文件不存在
		return err
	}
	if info, err := os.Stat(newPath); err == nil {
		if info.IsDir() {
			return nil
		}
		// 已存在文件 适配多版本
		newPath = genMultiVersionName(oldPath, newPath, newPathRoot)
	}
	return os.Rename(oldPath, newPath)
}

// ParseXml 解析xml文件
func ParseXml(xmlFilePath string, ptr interface{}) (err error) {
	f, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return
	}
	err = xml.Unmarshal(f, ptr)
	if err != nil {
		return
	}
	t := reflect.TypeOf(ptr)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("参数应该为结构体指针")
	}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Tag == "" {
			continue
		}
		if xmlTag, ok := v.Type().Field(i).Tag.Lookup("xml"); ok &&
			!strings.Contains(xmlTag, "omitempty") &&
			v.Type().Field(i).Type.Kind() == reflect.Ptr &&
			v.Field(i).IsNil() {
			err = errors.New("无法解析的xml数据")
			break
		}
	}
	return
}

func genMultiVersionName(oldPath, newPath, newPathRoot string) (s string) {
	var ext string
	if newPathRoot != "" && strings.Contains(newPath, newPathRoot) {
		ext = newPath[len(newPathRoot):]
		s = newPathRoot
	} else {
		ext = filepath.Ext(oldPath)
		s = newPath[:len(newPath)-len(ext)]
	}
	if specials := getMovieSpecialInfo(GetFileName(oldPath)); len(specials) > 0 {
		s += " - " + strings.Join(specials, ".")
	}
	for i := 1; i > 0; i++ {
		if _, err := os.Stat(s); err != nil {
			break
		}
		s = newPath + fmt.Sprintf(" (%d)", i)
	}
	return s + ext
}

func getMovieSpecialInfo(filePath string) (specials []string) {
	// https://en.wikipedia.org/wiki/Pirated_movie_release_types
	piratedMovieReleaseTypes := []string{"CAM\\-Rip", "CAM", "HDCAM", "TS", "HDTS", "TELESYNC", "PDVD", "PreDVDRip", "WP", "WORKPRINT", "TC", "HDTC", "TELECINE", "PPV", "PPVRip", "SCR", "SCREENER", "DVDSCR", "DVDSCREENER", "BDSCR", "WEBSCREENER", "DDC", "R5", "R5\\.LINE", "R5\\.AC3\\.5\\.1\\.HQ", "DVDRip", "DVDMux", "DVDR", "DVD\\-Full", "Full\\-Rip", "ISO rip", "lossless rip", "untouched rip", "DVD\\-5", "DVD\\-9", "DSR", "DSRip", "SATRip", "DTHRip", "DVBRip", "HDTV", "PDTV", "DTVRip", "TVRip", "HDTVRip", "VODRip", "VODR", "HC", "HD\\-Rip", "WEB\\-Cap", "WEBCAP", "WEB\\sCap", "HDRip", "WEB\\-DLRip", "WEBRip", "WEB\\sRip", "WEB\\-Rip", "WEBDL", "WEB\\sDL", "WEB\\-DL", "WEB", "WEBRip", "Blu\\-Ray", "BluRay", "BLURAY", "BDRip", "BRip", "BRRip", "BDR", "BD25", "BD50", "BD5", "BD9", "BDMV", "BDISO", "COMPLETE\\.BLURAY"}
	// https://en.wikipedia.org/wiki/Graphics_display_resolution
	resolutionList := []string{"240i", "288i", "480i", "576i", "480p", "576p", "720p", "1080i", "1080p", "2160p", "4320p", "nHD", "qHD", "HD", "HD+", "FHD", "DCI\\s2K", "QHD", "QHD+", "4K\\sUHD", "DCI\\s4K", "5K", "8K\\sUHD", "16K"}
	// https://en.wikipedia.org/wiki/Video_coding_format
	videoCodecList := []string{"divx", "dv", "h\\.120", "h\\.261", "h\\.262", "h\\.263", "h\\.264", "h\\.265", "h\\.266", "h120", "h261", "h262", "h263", "h264", "h265", "h266", "hevc", "mj2", "mpeg1", "mpeg2", "mpeg4", "mpegh", "mpeg\\-1", "mpeg\\-2", "mpeg\\-4", "mpeg\\-h", "vc\\-1", "vcc", "x264", "x265", "vc1"}

	list := [][]string{resolutionList, piratedMovieReleaseTypes, videoCodecList}
	for _, r := range list {
		regex := regexp.MustCompile("(?i)\\W(" + strings.Join(r, "|") + ")\\W")
		if specialMatch := regex.FindStringSubmatch(filePath); len(specialMatch) > 0 && specialMatch[1] != "" {
			specials = append(specials, specialMatch[1])
		}
	}
	return
}
