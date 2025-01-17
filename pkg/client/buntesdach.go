package client

import "net/url"

type BuntesdachClient struct {
	apiUrl *url.URL
	Resources
}

func NewBuntesdachClient(apiUrl string) (*BuntesdachClient, error) {
	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}

	return &BuntesdachClient{
		apiUrl:    parsedUrl,
		Resources: NewResources(parsedUrl),
	}, nil
}
