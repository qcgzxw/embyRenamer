package renamer

import (
	"os"
	"path/filepath"
	"strings"
)

type TvShow struct {
	rootPath   string
	nfoPath    string
	tvShowInfo TvShowInfo
	dirFormat  string
}

func (this *TvShow) GetOriginName() (originName string) {
	return this.tvShowInfo.Title
}
func (this *TvShow) GetEmbyDirName() (originName string) {
	if this.dirFormat == "" || !strings.Contains(this.dirFormat, "{") {
		return ""
	}
	replacer := this.nameReplacer()
	return replacer.Replace(this.dirFormat)
}
func (this *TvShow) GetEmbyTitleName() (originName string) {
	return
}
func (this *TvShow) Rename() {
	this.renameFile()
}
func (this *TvShow) nameReplacer() *strings.Replacer {
	if this.tvShowInfo.Originaltitle == "" {
		this.tvShowInfo.Originaltitle = this.tvShowInfo.Title
	}
	if strings.Contains(this.tvShowInfo.Originaltitle, string(os.PathSeparator)) {
		this.tvShowInfo.Originaltitle = strings.Replace(this.tvShowInfo.Originaltitle, string(os.PathSeparator), " ", -1)
	}
	if strings.Contains(this.tvShowInfo.Title, string(os.PathSeparator)) {
		this.tvShowInfo.Title = strings.Replace(this.tvShowInfo.Title, string(os.PathSeparator), " ", -1)
	}
	return strings.NewReplacer(
		"{originaltitle}", this.tvShowInfo.Originaltitle,
		"{title}", this.tvShowInfo.Title,
		"{year}", this.tvShowInfo.Year,
		"{imdbid}", this.tvShowInfo.ImdbID,
		"{tmdbid}", this.tvShowInfo.Tmdbid,
		"{tvdbid}", this.tvShowInfo.Tvdbid,
		"{releasedate}", this.tvShowInfo.Releasedate,
		"{id}", this.tvShowInfo.ID,
	)
}
func (this *TvShow) renameFile() {
	var embyDirName, embyDirPath string
	if embyDirName = this.GetEmbyDirName(); embyDirName == "" {
		return
	}
	embyDirPath = this.rootPath + string(os.PathSeparator) + embyDirName
	if stat, err := os.Stat(embyDirPath); err != nil || !stat.IsDir() {
		err = os.MkdirAll(embyDirPath, os.ModePerm)
		if err != nil {
			println(err.Error())
		}
	}
	files, paths, _ := FilePathWalkCurrentDir(filepath.Dir(this.nfoPath))
	embyTitleReplacer := strings.NewReplacer(
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"tvshow.nfo", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"tvshow.nfo",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"season", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"season",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"folder.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"folder.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"poster.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"poster.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"cover.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"cover.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"default.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"default.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"movie.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"movie.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"clearart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"clearart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"backdrop.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"backdrop.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"fanart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"fanart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"background.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"background.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"extrafanart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"extrafanart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"banner.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"banner.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"disc.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"disc.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"cdart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"cdart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"clearlogo.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"clearlogo.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"logo.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"logo.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"thumb.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"thumb.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"landscape.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"landscape.",
	)
	var episodeNfoFiles = make(map[string]byte)
	for path, _ := range files {
		if newPath := embyTitleReplacer.Replace(path); newPath != path {
			OsRename(path, newPath)
		} else if strings.ToLower(GetFileExt(path)) == ".nfo" {
			episodeNfoFiles[path] = 0x0
		}
	}
	for _, path := range paths {
		files, _, _ := FilePathWalkDir(path)
		for p, _ := range files {
			if strings.ToLower(GetFileExt(p)) == ".nfo" {
				episodeNfoFiles[p] = 0x0
			}
		}
	}

	totalSeasons := len(paths)
	for path, _ := range episodeNfoFiles {
		totalEpisodes := CountFileInDir(path, ".nfo")
		client := scanEpisodeNfo(this.rootPath, path, this.tvShowInfo, uint(totalSeasons), uint(totalEpisodes))
		if client == nil {
			continue
		}
		client.Rename()
	}
}
