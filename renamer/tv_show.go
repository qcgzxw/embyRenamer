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
	if strings.Contains(this.tvShowInfo.Originaltitle, "/") {
		this.tvShowInfo.Originaltitle = strings.Replace(this.tvShowInfo.Originaltitle, "/", " ", -1)
	}
	if strings.Contains(this.tvShowInfo.Title, "/") {
		this.tvShowInfo.Title = strings.Replace(this.tvShowInfo.Title, "/", " ", -1)
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
	embyDirPath = this.rootPath + "/" + embyDirName
	if stat, err := os.Stat(embyDirPath); err != nil || !stat.IsDir() {
		err = os.MkdirAll(embyDirPath, os.ModePerm)
		if err != nil {
			println(err.Error())
		}
	}
	files, paths, _ := FilePathWalkDir(filepath.Dir(this.nfoPath))
	embyTitleReplacer := strings.NewReplacer(
		filepath.Dir(this.nfoPath)+"/tvshow.nfo", this.rootPath+"/"+embyDirName+"/tvshow.nfo",
		filepath.Dir(this.nfoPath)+"/season", this.rootPath+"/"+embyDirName+"/season",
		filepath.Dir(this.nfoPath)+"/folder.", this.rootPath+"/"+embyDirName+"/folder.",
		filepath.Dir(this.nfoPath)+"/poster.", this.rootPath+"/"+embyDirName+"/poster.",
		filepath.Dir(this.nfoPath)+"/cover.", this.rootPath+"/"+embyDirName+"/cover.",
		filepath.Dir(this.nfoPath)+"/default.", this.rootPath+"/"+embyDirName+"/default.",
		filepath.Dir(this.nfoPath)+"/movie.", this.rootPath+"/"+embyDirName+"/movie.",
		filepath.Dir(this.nfoPath)+"/clearart.", this.rootPath+"/"+embyDirName+"/clearart.",
		filepath.Dir(this.nfoPath)+"/backdrop.", this.rootPath+"/"+embyDirName+"/backdrop.",
		filepath.Dir(this.nfoPath)+"/fanart.", this.rootPath+"/"+embyDirName+"/fanart.",
		filepath.Dir(this.nfoPath)+"/background.", this.rootPath+"/"+embyDirName+"/background.",
		filepath.Dir(this.nfoPath)+"/extrafanart.", this.rootPath+"/"+embyDirName+"/extrafanart.",
		filepath.Dir(this.nfoPath)+"/banner.", this.rootPath+"/"+embyDirName+"/banner.",
		filepath.Dir(this.nfoPath)+"/disc.", this.rootPath+"/"+embyDirName+"/disc.",
		filepath.Dir(this.nfoPath)+"/cdart.", this.rootPath+"/"+embyDirName+"/cdart.",
		filepath.Dir(this.nfoPath)+"/clearlogo.", this.rootPath+"/"+embyDirName+"/clearlogo.",
		filepath.Dir(this.nfoPath)+"/logo.", this.rootPath+"/"+embyDirName+"/logo.",
		filepath.Dir(this.nfoPath)+"/thumb.", this.rootPath+"/"+embyDirName+"/thumb.",
		filepath.Dir(this.nfoPath)+"/landscape.", this.rootPath+"/"+embyDirName+"/landscape.",
	)
	var episodeNfoFiles = make(map[string]os.FileInfo)
	for path, f := range files {
		if newPath := embyTitleReplacer.Replace(path); newPath != path {
			OsRename(path, newPath)
		} else if strings.ToLower(GetFileExt(path)) == ".nfo" {
			episodeNfoFiles[path] = f
		}
	}
	for _, path := range paths {
		files, _, _ = FilePathWalkDir(path)
		for p, f := range files {
			if strings.ToLower(GetFileExt(p)) == ".nfo" {
				episodeNfoFiles[p] = f
			}
		}
	}

	totalSeasons := len(paths)
	totalEpisodes := len(episodeNfoFiles)
	for path, _ := range episodeNfoFiles {
		client := scanEpisodeNfo(this.rootPath, path, this.tvShowInfo, uint(totalSeasons), uint(totalEpisodes))
		if client == nil {
			continue
		}
		client.Rename()
	}
}
