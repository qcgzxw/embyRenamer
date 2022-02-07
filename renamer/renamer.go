package renamer

import (
	"encoding/xml"
	"os"
	"strings"
)

func scanTvShowNfo(rootPath string, nfoPath string) (client Client) {
	if ext := GetFileExt(nfoPath); strings.ToLower(ext) != ".nfo" {
		return
	}
	if name := GetFileName(nfoPath); strings.ToLower(name) != "tvshow" {
		return
	}
	d, err := os.ReadFile(nfoPath)
	if err != nil {
		return
	}
	var tvShowInfo *TvShowInfo
	err = xml.Unmarshal(d, &tvShowInfo)
	if err != nil || tvShowInfo == nil {
		println(nfoPath)
		return
	}
	client = &TvShow{
		rootPath:   rootPath,
		nfoPath:    nfoPath,
		tvShowInfo: *tvShowInfo,
		dirFormat:  config["tvDirFormat"],
	}
	return
}
func scanEpisodeNfo(rootPath string, nfoPath string, tvShowInfo TvShowInfo, totalSeasons uint, totalEpisodes uint) (client Client) {
	if ext := GetFileExt(nfoPath); strings.ToLower(ext) != ".nfo" {
		return
	}
	if name := GetFileName(nfoPath); strings.ToLower(name) == "season" {
		return
	}
	d, err := os.ReadFile(nfoPath)
	if err != nil {
		return
	}
	var episodeInfo *EpisodeDetailsInfo
	err = xml.Unmarshal(d, &episodeInfo)
	if err != nil || episodeInfo == nil {
		return
	}
	client = &Episode{
		rootPath:      rootPath,
		nfoPath:       nfoPath,
		tvShowInfo:    tvShowInfo,
		episodeName:   tvShowInfo.Title,
		episodeInfo:   *episodeInfo,
		dirFormat:     config["tvDirFormat"] + "/" + config["episodeDirFormat"],
		titleFormat:   config["episodeTitleFormat"],
		totalSeasons:  totalSeasons,
		totalEpisodes: totalEpisodes,
	}
	return
}
func scanMovieNfo(rootPath string, nfoPath string) (client Client) {
	if ext := GetFileExt(nfoPath); strings.ToLower(ext) != ".nfo" {
		return
	}
	d, err := os.ReadFile(nfoPath)
	if err != nil {
		println(err.Error())
		return
	}
	var movieInfo *MovieInfo
	err = xml.Unmarshal(d, &movieInfo)
	if err != nil {
		return
	}
	client = &Movie{
		rootPath:    rootPath,
		nfoPath:     nfoPath,
		movieInfo:   *movieInfo,
		movieName:   GetFileName(nfoPath),
		dirFormat:   config["movieDirFormat"],
		titleFormat: config["movieTitleFormat"],
	}
	return
}

// Scan 扫描目录下的nfo文件并返回client实例
func Scan(rootPath, dirPath string) (clients []Client, err error) {
	var files = make(map[string]os.FileInfo)
	var nfoFiles = make(map[string]os.FileInfo)
	files, _, err = FilePathWalkDir(dirPath)
	if err != nil || len(files) == 0 {
		return
	}
	for path, file := range files {
		if ext := GetFileExt(path); ext == ".nfo" {
			nfoFiles[path] = file
		}
		if file.Name() == "tvshow.nfo" {
			if c := scanTvShowNfo(rootPath, path); c != nil {
				clients = append(clients, c)
			}
			return
		}
	}
	for path, _ := range nfoFiles {
		if c := scanMovieNfo(rootPath, path); c != nil {
			clients = append(clients, c)
		}
	}
	return
}

// DeepScan 深度扫描目录下的nfo文件并返回client实例
// rootPath: 根目录
// dirPath: 需要扫描的目录
// deep: 层数，小于0则无限制
func DeepScan(rootPath, dirPath string, deep int) (clients []Client, err error) {
	if _, err = os.Stat(rootPath); err != nil {
		if err = os.MkdirAll(rootPath, os.ModePerm); err != nil {
			return
		} else {
			err = nil
		}
	}
	if _, err = os.Stat(dirPath); err != nil {
		return
	}
	if deep == 0 {
		return
	}
	currentPathClient, err := Scan(rootPath, dirPath)
	if err != nil {
		println(err.Error())
		return
	}
	for _, c := range currentPathClient {
		if c != nil {
			clients = append(clients, c)
		}
	}
	_, dirPaths, err := FilePathWalkDir(dirPath)
	if err != nil || len(dirPaths) == 0 {
		return
	}
	for _, path := range dirPaths {
		c, e := DeepScan(rootPath, path, deep-1)
		if e == nil {
			clients = append(clients, c...)
		} else {
			println(e.Error())
		}
	}
	return
}
