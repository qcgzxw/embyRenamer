package renamer

import (
	"encoding/xml"
	"testing"
)

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
func TestParseXml(t *testing.T) {
	type movieInfo struct {
		XMLName       xml.Name `xml:"movie,omitempty"`
		Text          string   `xml:",chardata"`
		Plot          string   `xml:"plot,omitempty"`
		Outline       string   `xml:"outline,omitempty"`
		Lockdata      string   `xml:"lockdata,omitempty"`
		Dateadded     string   `xml:"dateadded,omitempty"`
		Title         *string  `xml:"title"`
		Originaltitle string   `xml:"originaltitle,omitempty"`
		Rating        string   `xml:"rating,omitempty"`
		Year          *string  `xml:"year"`
		Mpaa          string   `xml:"mpaa,omitempty"`
		Imdbid        *string  `xml:"imdbid"`
		Tmdbid        *string  `xml:"tmdbid,omitempty"`
		Premiered     string   `xml:"premiered,omitempty"`
		Releasedate   *string  `xml:"releasedate"`
		Runtime       string   `xml:"runtime,omitempty"`
		Country       *string  `xml:"country,omitempty"`
		Genre         []string `xml:"genre,omitempty"`
		Studio        []string `xml:"studio,omitempty"`
		Uniqueid      []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr,omitempty"`
		} `xml:"uniqueid,omitempty"`
		ID *string `xml:"id"`
	}
	var data = new(movieInfo)
	if err := ParseXml("/home/owen/go/src/embyRenamer/example/newData/电影/风中有朵雨做的云 (2018) [imdbid=tt8071308]/风中有朵雨做的云.nfo", data); err != nil {
		t.Fatal(err.Error())
	}
}
