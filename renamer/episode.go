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
	if strings.Contains(this.episodeInfo.Title, string(os.PathSeparator)) {
		this.episodeInfo.Title = strings.Replace(this.episodeInfo.Title, string(os.PathSeparator), " ", -1)
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
	embyDirPath = this.rootPath + string(os.PathSeparator) + embyDirName
	if stat, err := os.Stat(embyDirPath); err != nil || !stat.IsDir() {
		err = os.MkdirAll(embyDirPath, os.ModePerm)
		if err != nil {
			println(err.Error())
		}
	}
	files, _, _ := FilePathWalkCurrentDir(filepath.Dir(this.nfoPath))
	embyTitleReplacer := strings.NewReplacer(
		this.nfoPath[:len(this.nfoPath)-4]+".", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName+".",
		this.nfoPath[:len(this.nfoPath)-4]+"-", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName+"-",
		this.nfoPath[:len(this.nfoPath)-4]+"_", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName+"_",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"folder.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"folder.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"poster.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"poster.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"cover.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"cover.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"default.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"default.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"show.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"show.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"clearart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"clearart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"backdrop.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"backdrop.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"backdropX.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"backdropX.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"fanart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"fanart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"fanart-X.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"fanart-X.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"background.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"background.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"background-X.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"background-X.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"art.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"art.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"art-X.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"art-X.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"extrafanart", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"extrafanart",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"banner.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"banner.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"disc.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"disc.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"cdart.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"cdart.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"logo.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"logo.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"thumb.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"thumb.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"landscape.", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"landscape.",
		filepath.Dir(this.nfoPath)+string(os.PathSeparator)+"season", this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"season",
	)
	for path, _ := range files {
		if newPath := embyTitleReplacer.Replace(path); newPath != path {
			OsRename(path, newPath, this.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName)
		}
	}
}
