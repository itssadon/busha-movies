package cache

import "github.com/itssadon/busha-movies/collections"

type MoviesCache interface {
	Set(key string, value *collections.MovieCollection)
	Get(key string) *collections.MovieCollection
}
