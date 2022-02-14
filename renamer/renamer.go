package renamer

import (
	"os"
	"strings"
)

var config Config
var invalidNfoPath = make(map[string]byte)

func scanTvShowNfo(rootPath string, nfoPath string) (client Client) {
	if ext := GetFileExt(nfoPath); strings.ToLower(ext) != ".nfo" || strings.ToLower(GetFileName(nfoPath)) != "tvshow" {
		return
	}
	tvShowInfo := new(TvShowInfo)
	err := ParseXml(nfoPath, tvShowInfo)
	if err != nil {
		if _, ok := invalidNfoPath[nfoPath]; !ok {
			invalidNfoPath[nfoPath] = 0x0
		}
		return
	}
	client = &TvShow{
		rootPath:   rootPath,
		nfoPath:    nfoPath,
		tvShowInfo: *tvShowInfo,
		dirFormat:  config.TvDirFormat,
	}
	return
}
func scanEpisodeNfo(rootPath string, nfoPath string, tvShowInfo TvShowInfo, totalSeasons uint, totalEpisodes uint) (client Client) {
	if ext := GetFileExt(nfoPath); strings.ToLower(ext) != ".nfo" || strings.ToLower(GetFileName(nfoPath)) == "season" {
		return
	}
	episodeInfo := new(EpisodeDetailsInfo)
	err := ParseXml(nfoPath, episodeInfo)
	if err != nil || episodeInfo == nil {
		if _, ok := invalidNfoPath[nfoPath]; !ok {
			invalidNfoPath[nfoPath] = 0x0
		}
		return
	}
	client = &Episode{
		rootPath:      rootPath,
		nfoPath:       nfoPath,
		tvShowInfo:    tvShowInfo,
		episodeName:   *tvShowInfo.Title,
		episodeInfo:   *episodeInfo,
		dirFormat:     config.TvDirFormat + string(os.PathSeparator) + config.EpisodeDirFormat,
		titleFormat:   config.EpisodeTitleFormat,
		totalSeasons:  totalSeasons,
		totalEpisodes: totalEpisodes,
	}
	return
}
func scanMovieNfo(rootPath string, nfoPath string) (client Client) {
	if ext := GetFileExt(nfoPath); strings.ToLower(ext) != ".nfo" || strings.ToLower(GetFileName(nfoPath)) == "season" {
		return
	}
	movieInfo := new(MovieInfo)
	err := ParseXml(nfoPath, movieInfo)
	if err != nil {
		if _, ok := invalidNfoPath[nfoPath]; err.Error() == "无法解析的xml数据" && !ok {
			invalidNfoPath[nfoPath] = 0x0
		}
		return
	}
	client = &Movie{
		rootPath:    rootPath,
		nfoPath:     nfoPath,
		movieInfo:   *movieInfo,
		movieName:   GetFileName(nfoPath),
		dirFormat:   config.MovieDirFormat,
		titleFormat: config.MovieTitleFormat,
	}
	return
}

// Scan 扫描目录下的nfo文件并返回client实例
func Scan(rootPath, dirPath string) (clients []Client, err error) {
	var files = make(map[string]os.DirEntry)
	var nfoFiles = make(map[string]os.DirEntry)
	var isMoviePath = true
	files, _, err = FilePathWalkCurrentDir(dirPath)
	if err != nil || len(files) == 0 {
		return
	}
	for path, file := range files {
		if strings.ToLower(file.Name()) == "season.nfo" {
			continue
		}
		if strings.ToLower(file.Name()) == "tvshow.nfo" {
			isMoviePath = false
			if c := scanTvShowNfo(rootPath, path); c != nil {
				clients = append(clients, c)
			}
			continue
		}
		if ext := GetFileExt(path); ext == ".nfo" {
			nfoFiles[path] = file
		}
	}
	if isMoviePath {
		for path, _ := range nfoFiles {
			if c := scanMovieNfo(rootPath, path); c != nil {
				clients = append(clients, c)
			}
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
	_, dirPaths, err := FilePathWalkCurrentDir(dirPath)
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

func InvalidNfoPath() map[string]byte {
	return invalidNfoPath
}

func LoadConfig(c *Config) {
	if c != nil {
		config = *c
	} else {
		// load default config
		config = defaultConfig
	}
}
