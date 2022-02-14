package renamer

import (
	"os"
	"path/filepath"
	"strings"
)

type Movie struct {
	rootPath    string
	nfoPath     string
	movieInfo   MovieInfo
	movieName   string
	dirFormat   string
	titleFormat string
}

func (m *Movie) GetOriginName() (originName string) {
	return m.movieName
}
func (m *Movie) GetEmbyDirName() string {
	if m.dirFormat == "" || !strings.Contains(m.dirFormat, "{") {
		return ""
	}
	replacer := m.nameReplacer()
	return replacer.Replace(m.dirFormat)
}
func (m *Movie) GetEmbyTitleName() string {
	if m.titleFormat == "" || !strings.Contains(m.titleFormat, "{") {
		return ""
	}
	replacer := m.nameReplacer()
	return replacer.Replace(m.titleFormat)
}
func (m *Movie) Rename() {
	m.renameFile()
}
func (m *Movie) nameReplacer() *strings.Replacer {
	if m.movieInfo.Originaltitle == "" {
		m.movieInfo.Originaltitle = *m.movieInfo.Title
	}
	if strings.Contains(m.movieInfo.Originaltitle, string(os.PathSeparator)) {
		m.movieInfo.Originaltitle = strings.Replace(m.movieInfo.Originaltitle, string(os.PathSeparator), " ", -1)
	}
	if strings.Contains(*m.movieInfo.Title, string(os.PathSeparator)) {
		if tmp := strings.Replace(*m.movieInfo.Title, string(os.PathSeparator), " ", -1); tmp != "" {
			m.movieInfo.Title = &tmp
		}
	}
	return strings.NewReplacer(
		"{originaltitle}", m.movieInfo.Originaltitle,
		"{title}", *m.movieInfo.Title,
		"{year}", *m.movieInfo.Year,
		"{imdbid}", *m.movieInfo.Imdbid,
		"{tmdbid}", m.movieInfo.Tmdbid,
		"{releasedate}", m.movieInfo.Releasedate,
		"{country}", m.movieInfo.Country,
		"{id}", *m.movieInfo.ID,
	)
}

func (m *Movie) renameFile() {
	var embyTitleName, embyDirName, embyDirPath string
	if embyTitleName = m.GetEmbyTitleName(); embyTitleName == "" {
		return
	}
	if embyDirName = m.GetEmbyDirName(); embyDirName == "" {
		return
	}
	embyDirPath = m.rootPath + string(os.PathSeparator) + embyDirName
	if stat, err := os.Stat(embyDirPath); err != nil || !stat.IsDir() {
		err = os.MkdirAll(embyDirPath, os.ModePerm)
		if err != nil {
			println(err.Error())
		}
	}
	files, paths, _ := FilePathWalkCurrentDir(filepath.Dir(m.nfoPath))
	embyTitleReplacer := strings.NewReplacer(
		m.nfoPath[:len(m.nfoPath)-4]+".", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName+".",
		m.nfoPath[:len(m.nfoPath)-4]+"-", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName+"-",
		m.nfoPath[:len(m.nfoPath)-4]+"_", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName+"_",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"sample", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"sample",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"folder.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"folder.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"poster.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"poster.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"cover.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"cover.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"default.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"default.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"movie.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"movie.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"clearart.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"clearart.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"backdrop.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"backdrop.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"backdropX.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"backdropX.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"fanart.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"fanart.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"fanart-X.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"fanart-X.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"background.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"background.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"background-X.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"background-X.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"art.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"art.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"art-X.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"art-X.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"extrafanart.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"extrafanart.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"banner.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"banner.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"disc.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"disc.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"cdart.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"cdart.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"clearlogo.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"clearlogo.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"logo.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"logo.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"thumb.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"thumb.",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"landscape.", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"landscape.",
	)
	for path, _ := range files {
		if newPath := embyTitleReplacer.Replace(path); newPath != path {
			OsRename(path, newPath, m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName)
		}
	}
	if m.rootPath == filepath.Dir(m.nfoPath) {
		return
	}
	embyDirReplacer := strings.NewReplacer(
		m.nfoPath[:len(m.nfoPath)-4], m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName,
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"Sample", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"Sample",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"BDMV", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"BDMV",
		filepath.Dir(m.nfoPath)+string(os.PathSeparator)+"CERTIFICATE", m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+"CERTIFICATE",
	)
	for _, path := range paths {
		if newPath := embyDirReplacer.Replace(path); newPath != path {
			OsRename(path, newPath, m.rootPath+string(os.PathSeparator)+embyDirName+string(os.PathSeparator)+embyTitleName)
		}
	}
	return
}
func (m *Movie) moveFile() {
	return
}
