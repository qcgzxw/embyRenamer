package renamer

import "testing"

func TestDeepScan(t *testing.T) {
	clients, err := DeepScan("/home/owen/下载/test/example/data/电影New", "/home/owen/下载/test/example/data/电影", -1)
	//clients, err := DeepScan("/home/owen/下载/test/example/data/MovieNew", "/home/owen/下载/test/example/data/Movie", 2)
	if err != nil {
		println(err.Error())
		return
	}
	for _, client := range clients {
		client.Rename()
	}
}
