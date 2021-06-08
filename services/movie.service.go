package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/itssadon/busha-movies/collections"
)

const swapiBaseUrl = "https://swapi.dev/api"

func GetSwapiFilms() (*collections.MovieCollection, error) {
	result := &collections.MovieCollection{}

	response, err := http.Get(fmt.Sprintf("%s%s", swapiBaseUrl, "/films/"))
	if err != nil {
		return nil, err
	}

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, readError
	}

	closeErr := response.Body.Close()
	if closeErr != nil {
		return nil, closeErr
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
