package client

import (
	"fmt"
	v1 "github.com/kyzrfranz/buntesdach/api/v1"
	"net/url"
)

const ResourceNameCommittees = "committees"

type Committees interface {
	List() []v1.CommitteeListEntry
	Get(id string) *v1.CommitteeListEntry
	Detail(id string) *v1.CommitteeDetails
}

type committees struct {
	url *url.URL
}

func NewCommittees(apiUrl *url.URL) Committees {
	return &committees{
		url: apiUrl,
	}
}

func (c committees) List() []v1.CommitteeListEntry {
	c.url.Path = fmt.Sprintf("/%s", ResourceNameCommittees)
	list, err := executeRequest[[]v1.CommitteeListEntry](c.url)
	if err != nil {
		return nil
	}
	return list
}

func (c committees) Get(id string) *v1.CommitteeListEntry {
	c.url.Path = fmt.Sprintf("/%s/%s", ResourceNameCommittees, id)
	committee, err := executeRequest[v1.CommitteeListEntry](c.url)
	if err != nil {
		return nil
	}
	return &committee
}

func (c committees) Detail(id string) *v1.CommitteeDetails {
	c.url.Path = fmt.Sprintf("/%s/%s/detail", ResourceNameCommittees, id)
	committee, err := executeRequest[v1.CommitteeDetails](c.url)
	if err != nil {
		return nil
	}
	return &committee
}
