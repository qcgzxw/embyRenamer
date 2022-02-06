package renamer

type Episode struct {
	rootPath    string
	nfoPath     string
	tvShowInfo  TvShowInfo
	tvShowName  string
	titleFormat string
}

func (m *Episode) GetOriginName() (originName string) {
	return m.tvShowName
}
func (m *Episode) GetEmbyDirName() (originName string) {
	return
}
func (m *Episode) GetEmbyTitleName() (originName string) {
	return
}
func (m *Episode) Rename() {
	return
}
