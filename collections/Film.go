package collections

type FilmCollection struct {
	Results []Film `json:"results"`
}

type Film struct {
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	ReleaseDate  string `json:"release_date"`
}
