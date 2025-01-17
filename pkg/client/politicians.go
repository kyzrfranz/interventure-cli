package client

import (
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"net/url"
)

const ResourceNamePoliticians = "politicians"

type Politicians interface {
	List() []v1.PersonListEntry
	Get(id string) *v1.PersonListEntry
	Bio(id string) *v1.Politician
}

type politicians struct {
	url *url.URL
}

func NewPoliticians(apiUrl *url.URL) Politicians {
	return &politicians{
		url: apiUrl,
	}
}

func (p politicians) List() []v1.PersonListEntry {
	p.url.Path = fmt.Sprintf("/%s", ResourceNamePoliticians)
	list, err := executeRequest[[]v1.PersonListEntry](p.url)
	if err != nil {
		return nil
	}
	return list
}

func (p politicians) Get(id string) *v1.PersonListEntry {
	p.url.Path = fmt.Sprintf("/%s/%s", ResourceNamePoliticians, id)
	person, err := executeRequest[v1.PersonListEntry](p.url)
	if err != nil {
		return nil
	}
	return &person
}

func (p politicians) Bio(id string) *v1.Politician {
	p.url.Path = fmt.Sprintf("/%s/%s/bio", ResourceNamePoliticians, id)
	person, err := executeRequest[v1.Politician](p.url)
	if err != nil {
		return nil
	}
	return &person
}
