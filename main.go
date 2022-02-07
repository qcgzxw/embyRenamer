package main

import (
	"embyRenamer/renamer"
	"encoding/json"
	"os"
)

func main() {
	var config *renamer.Config
	if b, err := os.ReadFile("config.json"); err == nil && len(b) > 0 {
		err = json.Unmarshal(b, &config)
		if err == nil {
			renamer.LoadConfig(config)
		}
	}
	if config == nil {
		panic("无法获取正确的配置")
	}
	if err := config.Check(); err != nil {
		panic(err)
	}
	if config.MovieDirPath != "" {
		// 电影
		if clients, err := renamer.DeepScan(config.MovieRootPath, config.MovieDirPath, -1); err == nil {
			for _, client := range clients {
				client.Rename()
			}
		} else {
			panic(err)
		}
	}
	if config.MovieDirPath != "" {
		// 电视剧
		if clients, err := renamer.DeepScan(config.TvRootPath, config.TvDirPath, -1); err == nil {
			for _, client := range clients {
				client.Rename()
			}
		} else {
			panic(err)
		}
	}
}
