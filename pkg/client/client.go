package client

import (
	"encoding/json"
	"github.com/kyzrfranz/interventure-cli/pkg/http"
	"net/url"
)

type Resources interface {
	Politicians() Politicians
	Committees() Committees
}

type resources struct {
	politicians Politicians
	committees  Committees
}

func NewResources(apiUrl *url.URL) Resources {
	return &resources{
		politicians: NewPoliticians(apiUrl),
		committees:  NewCommittees(apiUrl),
	}
}

func (r resources) Politicians() Politicians {
	return r.politicians
}

func (r resources) Committees() Committees {
	return r.committees
}

func executeRequest[T any](url *url.URL) (T, error) {
	data, err := http.FetchUrl(url)
	if err != nil {
		var zero T
		return zero, err
	}

	var result T
	json.Unmarshal(data, &result)
	return result, nil
}
