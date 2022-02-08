package renamer

import (
	"os"
	"path/filepath"
	"strings"
)

type Episode struct {
	rootPath      string
	nfoPath       string
	tvShowInfo    TvShowInfo
	episodeInfo   EpisodeDetailsInfo
	episodeName   string
	dirFormat     string
	titleFormat   string
	totalSeasons  uint
	totalEpisodes uint
}

func (this *Episode) GetOriginName() (originName string) {
	return this.episodeName
}
func (this *Episode) GetEmbyDirName() string {
	if this.dirFormat == "" || !strings.Contains(this.dirFormat, "{") {
		return ""
	}
	replacer := this.nameReplacer()
	return replacer.Replace(this.dirFormat)
}
func (this *Episode) GetEmbyTitleName() string {
	if this.titleFormat == "" || !strings.Contains(this.titleFormat, "{") {
		return ""
	}
	replacer := this.nameReplacer()
	return replacer.Replace(this.titleFormat)
}
func (this *Episode) Rename() {
	this.renameFile()
}
func (this *Episode) nameReplacer() *strings.Replacer {
	if strings.Contains(this.episodeInfo.Title, "/") {
		this.episodeInfo.Title = strings.Replace(this.episodeInfo.Title, "/", " ", -1)
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
		"{season}", this.episodeInfo.Season,
		"{SEASON}", GetNumStr(this.totalSeasons, this.episodeInfo.Season),
		"{episode}", this.episodeInfo.Episode,
		"{EPISODE}", GetNumStr(this.totalEpisodes, this.episodeInfo.Episode),
	)
}

func (this *Episode) renameFile() {
	var embyTitleName, embyDirName, embyDirPath string
	if embyTitleName = this.GetEmbyTitleName(); embyTitleName == "" {
		return
	}
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
	files, _, _ := FilePathWalkDir(filepath.Dir(this.nfoPath))
	embyTitleReplacer := strings.NewReplacer(
		this.nfoPath[:len(this.nfoPath)-4], this.rootPath+"/"+embyDirName+"/"+embyTitleName,
		filepath.Dir(this.nfoPath)+"/"+"season.nfo", this.rootPath+"/"+embyDirName+"/"+"season.nfo",
		filepath.Dir(this.nfoPath)+"/"+"season-specials", this.rootPath+"/"+embyDirName+"/"+"season-specials",
	)
	for path, _ := range files {
		if newPath := embyTitleReplacer.Replace(path); newPath != path {
			OsRename(path, newPath)
		}
	}
}
