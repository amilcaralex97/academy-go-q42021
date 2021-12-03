package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"go-project/app/domain"
)

type ApiRepo struct {
	url string
}

func NewApiRepo(url string) ApiRepo {
	return ApiRepo{
		url: url,
	}
}

//FetchCharacters gets characters from an API
func (ar ApiRepo) FetchCharacters() (domain.Characters, error) {
	response, err := http.Get(ar.url)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	responseObject := Response{}
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Results, nil
}
