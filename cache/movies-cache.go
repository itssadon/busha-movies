package cache

import "github.com/itssadon/busha-movies/collections"

type MoviesCache interface {
	Set(key string, value *collections.FilmCollection)
	Get(key string) *collections.FilmCollection
}
