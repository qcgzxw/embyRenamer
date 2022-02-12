package main

import (
	"embyRenamer/renamer"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	if config.MovieRename && config.MovieDirPath != "" {
		// 电影
		if clients, err := renamer.DeepScan(strings.TrimRight(config.MovieRootPath, string(os.PathSeparator)), strings.TrimRight(config.MovieDirPath, string(os.PathSeparator)), -1); err == nil {
			for _, client := range clients {
				client.Rename()
			}
		} else {
			panic(err)
		}
	}
	if config.TvRename && config.TvDirPath != "" {
		// 电视剧
		if clients, err := renamer.DeepScan(strings.TrimRight(config.TvRootPath, string(os.PathSeparator)), strings.TrimRight(config.TvDirPath, string(os.PathSeparator)), -1); err == nil {
			for _, client := range clients {
				client.Rename()
			}
		} else {
			panic(err)
		}
	}
	endTime := time.Now().UnixNano() / 1e6
	if invalidNfoPath := renamer.InvalidNfoPath(); len(invalidNfoPath) > 0 {
		println("无法识别的nfo文件列表：\r\n")
		for path := range invalidNfoPath {
			println("  ", path)
		}
	}
	println()
	fmt.Printf("总花费时间：%dms\r\n", endTime-startTime)
}
