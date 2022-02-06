package renamer

import "encoding/xml"

type MovieInfo struct {
	XMLName       xml.Name `xml:"movie,omitempty"`
	Text          string   `xml:",chardata"`
	Plot          string   `xml:"plot,omitempty"`
	Outline       string   `xml:"outline,omitempty"`
	Lockdata      string   `xml:"lockdata,omitempty"`
	Dateadded     string   `xml:"dateadded,omitempty"`
	Title         string   `xml:"title,omitempty"`
	Originaltitle string   `xml:"originaltitle,omitempty"`
	Rating        string   `xml:"rating,omitempty"`
	Year          string   `xml:"year,omitempty"`
	Mpaa          string   `xml:"mpaa,omitempty"`
	Imdbid        string   `xml:"imdbid,omitempty"`
	Tmdbid        string   `xml:"tmdbid,omitempty"`
	Premiered     string   `xml:"premiered,omitempty"`
	Releasedate   string   `xml:"releasedate,omitempty"`
	Runtime       string   `xml:"runtime,omitempty"`
	Country       string   `xml:"country,omitempty"`
	Genre         []string `xml:"genre,omitempty"`
	Studio        []string `xml:"studio,omitempty"`
	Uniqueid      []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr,omitempty"`
	} `xml:"uniqueid,omitempty"`
	ID string `xml:"id,omitempty"`
}

type TvShowInfo struct {
	XMLName       xml.Name `xml:"tvshow,omitempty"`
	Text          string   `xml:",chardata"`
	Plot          string   `xml:"plot,omitempty"`
	Outline       string   `xml:"outline,omitempty"`
	Lockdata      string   `xml:"lockdata,omitempty"`
	Dateadded     string   `xml:"dateadded,omitempty"`
	Title         string   `xml:"title,omitempty"`
	Originaltitle string   `xml:"originaltitle,omitempty"`
	Trailer       []string `xml:"trailer,omitempty"`
	Rating        string   `xml:"rating,omitempty"`
	Year          string   `xml:"year,omitempty"`
	Mpaa          string   `xml:"mpaa,omitempty"`
	ImdbID        string   `xml:"imdb_id,omitempty"`
	Tmdbid        string   `xml:"tmdbid,omitempty"`
	Premiered     string   `xml:"premiered,omitempty"`
	Releasedate   string   `xml:"releasedate,omitempty"`
	Runtime       string   `xml:"runtime,omitempty"`
	Genre         []string `xml:"genre,omitempty"`
	Studio        []string `xml:"studio,omitempty"`
	Uniqueid      []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr,omitempty"`
	} `xml:"uniqueid,omitempty"`
	Tvdbid       string `xml:"tvdbid,omitempty"`
	ID           string `xml:"id,omitempty"`
	Season       string `xml:"season,omitempty"`
	Episode      string `xml:"episode,omitempty"`
	Displayorder string `xml:"displayorder,omitempty"`
	Status       string `xml:"status,omitempty"`
}

type SeasonInfo struct {
	XMLName      xml.Name `xml:"season,omitempty"`
	Text         string   `xml:",chardata"`
	Plot         string   `xml:"plot,omitempty"`
	Outline      string   `xml:"outline,omitempty"`
	Lockdata     string   `xml:"lockdata,omitempty"`
	Dateadded    string   `xml:"dateadded,omitempty"`
	Title        string   `xml:"title,omitempty"`
	Year         string   `xml:"year,omitempty"`
	Premiered    string   `xml:"premiered,omitempty"`
	Releasedate  string   `xml:"releasedate,omitempty"`
	Seasonnumber string   `xml:"seasonnumber,omitempty"`
}

type EpisodeDetailsInfo struct {
	XMLName   xml.Name `xml:"episodedetails,omitempty"`
	Text      string   `xml:",chardata"`
	Plot      string   `xml:"plot,omitempty"`
	Outline   string   `xml:"outline,omitempty"`
	Lockdata  string   `xml:"lockdata,omitempty"`
	Dateadded string   `xml:"dateadded,omitempty"`
	Title     string   `xml:"title,omitempty"`
	Rating    string   `xml:"rating,omitempty"`
	Year      string   `xml:"year,omitempty"`
	Runtime   string   `xml:"runtime,omitempty"`
	Episode   string   `xml:"episode,omitempty"`
	Season    string   `xml:"season,omitempty"`
	Aired     string   `xml:"aired,omitempty"`
}

var config = map[string]string{
	"movieDirPath":       "example/data/Movie/",
	"movieDirFormat":     "{originaltitle} ({year}) [imdbid={imdbid}]",
	"movieTitleFormat":   "{originaltitle}",
	"tvDirPath":          "example/data/TV/",
	"tvTitleFormat":      "",
	"seasonTitleFormat":  "",
	"episodeTitleFormat": "",
}
