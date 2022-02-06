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
	if m.titleFormat == "" || !strings.Contains(m.titleFormat, "{") {
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
		m.movieInfo.Originaltitle = m.movieInfo.Title
	}
	if strings.Contains(m.movieInfo.Originaltitle, "/") {
		m.movieInfo.Originaltitle = strings.Replace(m.movieInfo.Originaltitle, "/", " ", -1)
	}
	if strings.Contains(m.movieInfo.Title, "/") {
		m.movieInfo.Title = strings.Replace(m.movieInfo.Title, "/", " ", -1)
	}
	return strings.NewReplacer(
		"{originaltitle}", m.movieInfo.Originaltitle,
		"{title}", m.movieInfo.Title,
		"{year}", m.movieInfo.Year,
		"{imdbid}", m.movieInfo.Imdbid,
		"{tmdbid}", m.movieInfo.Tmdbid,
		"{releasedate}", m.movieInfo.Releasedate,
		"{country}", m.movieInfo.Country,
		"{id}", m.movieInfo.ID,
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
	embyDirPath = m.rootPath + "/" + embyDirName
	if stat, err := os.Stat(embyDirPath); err != nil || !stat.IsDir() {
		err = os.MkdirAll(embyDirPath, os.ModePerm)
		if err != nil {
			println(err.Error())
		}
	}
	files, paths, _ := FilePathWalkDir(filepath.Dir(m.nfoPath))
	embyTitleReplacer := strings.NewReplacer(
		strings.TrimRight(m.nfoPath, ".nfo"), m.rootPath+"/"+embyDirName+"/"+embyTitleName,
		filepath.Dir(m.nfoPath)+"/"+"folder.", m.rootPath+"/"+embyDirName+"/"+"folder.",
		filepath.Dir(m.nfoPath)+"/"+"poster.", m.rootPath+"/"+embyDirName+"/"+"poster.",
		filepath.Dir(m.nfoPath)+"/"+"cover.", m.rootPath+"/"+embyDirName+"/"+"cover.",
		filepath.Dir(m.nfoPath)+"/"+"default.", m.rootPath+"/"+embyDirName+"/"+"default.",
		filepath.Dir(m.nfoPath)+"/"+"movie.", m.rootPath+"/"+embyDirName+"/"+"movie.",
		filepath.Dir(m.nfoPath)+"/"+"clearart.", m.rootPath+"/"+embyDirName+"/"+"clearart.",
		filepath.Dir(m.nfoPath)+"/"+"backdrop.", m.rootPath+"/"+embyDirName+"/"+"backdrop.",
		filepath.Dir(m.nfoPath)+"/"+"fanart.", m.rootPath+"/"+embyDirName+"/"+"fanart.",
		filepath.Dir(m.nfoPath)+"/"+"background.", m.rootPath+"/"+embyDirName+"/"+"background.",
		filepath.Dir(m.nfoPath)+"/"+"extrafanart.", m.rootPath+"/"+embyDirName+"/"+"extrafanart.",
		filepath.Dir(m.nfoPath)+"/"+"banner.", m.rootPath+"/"+embyDirName+"/"+"banner.",
		filepath.Dir(m.nfoPath)+"/"+"disc.", m.rootPath+"/"+embyDirName+"/"+"disc.",
		filepath.Dir(m.nfoPath)+"/"+"cdart.", m.rootPath+"/"+embyDirName+"/"+"cdart.",
		filepath.Dir(m.nfoPath)+"/"+"clearlogo.", m.rootPath+"/"+embyDirName+"/"+"clearlogo .",
		filepath.Dir(m.nfoPath)+"/"+"logo.", m.rootPath+"/"+embyDirName+"/"+"logo.",
		filepath.Dir(m.nfoPath)+"/"+"thumb.", m.rootPath+"/"+embyDirName+"/"+"thumb.",
		filepath.Dir(m.nfoPath)+"/"+"landscape.", m.rootPath+"/"+embyDirName+"/"+"landscape.",
	)
	for path, _ := range files {
		if newPath := embyTitleReplacer.Replace(path); newPath != path {
			err := os.Rename(path, newPath)
			if err != nil {
				println(err.Error())
			}
		}
	}
	if m.rootPath == filepath.Dir(m.nfoPath) {
		return
	}
	embyDirReplacer := strings.NewReplacer(
		strings.TrimRight(m.nfoPath, ".nfo"), m.rootPath+"/"+embyDirName+"/"+embyTitleName,
		filepath.Dir(m.nfoPath)+"/"+"BDMV", m.rootPath+"/"+embyDirName+"/"+"BDMV",
		filepath.Dir(m.nfoPath)+"/"+"CERTIFICATE", m.rootPath+"/"+embyDirName+"/"+"CERTIFICATE",
	)
	for _, path := range paths {
		if newPath := embyDirReplacer.Replace(path); newPath != path {
			err := os.Rename(path, newPath)
			if err != nil {
				println(err.Error())
			}
		}
	}
	return
}
func (m *Movie) moveFile() {
	return
}
