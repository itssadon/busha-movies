package collections

type MovieCollection struct {
	Results []Movie `json:"results"`
}

type Movie struct {
	EpisodeId    int    `json:"episode_id"`
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	ReleaseDate  string `json:"release_date"`
}
