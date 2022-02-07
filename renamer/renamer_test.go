package renamer

import "testing"

func TestDeepScanMovie(t *testing.T) {
	LoadConfig(nil)
	clients, err := DeepScan(config.MovieRootPath, config.MovieDirPath, -1)
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, client := range clients {
		client.Rename()
	}
}

func TestDeepScanTV(t *testing.T) {
	LoadConfig(nil)
	clients, err := DeepScan(config.TvRootPath, config.TvDirPath, -1)
	if err != nil {
		t.Fatal(err.Error())
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
