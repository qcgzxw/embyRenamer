package main

import (
	"embyRenamer/renamer"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixNano() / 1e6
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
	endTime := time.Now().UnixNano() / 1e6
	if invalidNfoPath := renamer.InvalidNfoPath(); len(invalidNfoPath) > 0 {
		println("无法识别的nfo文件列表：")
		for path, _ := range invalidNfoPath {
			println(path)
		}
	}
	fmt.Printf("总花费时间：%dms\n", endTime-startTime)
}
