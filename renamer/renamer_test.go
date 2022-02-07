package renamer

import "testing"

func TestDeepScanMovie(t *testing.T) {
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

func TestDeepScanTV(t *testing.T) {
	clients, err := DeepScan("/home/owen/下载/test/example/data/电视剧New", "/home/owen/下载/test/example/data/电视剧", -1)
	//clients, err := DeepScan("/home/owen/下载/test/example/data/MovieNew", "/home/owen/下载/test/example/data/Movie", 2)
	if err != nil {
		println(err.Error())
		return
	}
	for _, client := range clients {
		client.Rename()
	}
}

func TestGetNumStr(t *testing.T) {
	if GetNumStr(1, "3") != "03" {
		t.Fatal("Error")
	}
	if GetNumStr(1000, "3") != "0003" {
		t.Fatal("Error")
	}
	if GetNumStr(10, "34") != "34" {
		t.Fatal("Error")
	}
	if GetNumStr(100, "344") != "344" {
		t.Fatal("Error")
	}
}
