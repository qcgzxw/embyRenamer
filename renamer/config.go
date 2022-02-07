package renamer

import "errors"

func (this *Config) Check() (err error) {
	if this.TvDirPath == "" && this.TvRootPath == "" && this.MovieDirPath == "" && this.MovieRootPath == "" {
		err = errors.New("配置有误")
	}
	if this.MovieDirPath != "" {
		if this.MovieRootPath == "" || this.MovieDirFormat == "" || this.MovieTitleFormat == "" {
			err = errors.New("配置有误")
		}
	}
	if this.TvDirPath != "" {
		if this.TvRootPath == "" || this.TvDirFormat == "" || this.EpisodeDirFormat == "" || this.EpisodeTitleFormat == "" {
			err = errors.New("配置有误")
		}
	}
	return
}
