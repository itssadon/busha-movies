package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/itssadon/busha-movies/cache"
	"github.com/itssadon/busha-movies/collections"
	"github.com/itssadon/busha-movies/services"
)

var (
	moviesCache cache.MoviesCache = cache.NewRedisCache("redis-15737.c52.us-east-1-4.ec2.cloud.redislabs.com:15737", 0, 60*60)
)

// ByReleaseDate implements sort.Interface based on the Age field.
type ByReleaseDate []collections.Film

func (a ByReleaseDate) Len() int           { return len(a) }
func (a ByReleaseDate) Less(i, j int) bool { return a[i].ReleaseDate < a[j].ReleaseDate }
func (a ByReleaseDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var movieList *collections.FilmCollection = moviesCache.Get("movies")
	if movieList == nil {
		movieList, err := services.GetSwapiFilms()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, err = res.Write([]byte(`{"status": false, "message": "Internal server error fetching movie list"`))
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		sort.Sort(sort.Reverse(ByReleaseDate(movieList.Results)))
		moviesCache.Set("movies", movieList)
		moviesJson, err := json.Marshal(movieList.Results)
		if err != nil {
			_, err = res.Write([]byte(`{"status": false, "message": "Error encoding movie list"`))
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		res.WriteHeader(http.StatusOK)
		_, err = res.Write(moviesJson)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		moviesJson, err := json.Marshal(movieList.Results)
		if err != nil {
			_, err = res.Write([]byte(`{"status": false, "message": "Error encoding movie list"`))
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		res.WriteHeader(http.StatusOK)
		_, err = res.Write(moviesJson)
		if err != nil {
			log.Fatal(err)
		}
	}
}
