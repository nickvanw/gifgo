// Package gifgo implements the GIPHY API in Go
package gifgo

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	apiHost         = "api.giphy.com"
	apiPath         = "v1/gifs/"
	stickersAPIPath = "v1/stickers/"
	apiKey          = "dc6zaTOxFJmzC"
)

var (
	searchPath, _    = url.Parse("search")
	translatePath, _ = url.Parse("translate")
	randomPath, _    = url.Parse("random")
	trendingPath, _  = url.Parse("trending")
	idsPath, _       = url.Parse("../gifs")
)

// Client represents a giphy API client
type Client struct {
	Stickers StickerClient

	client  *http.Client
	baseURL *url.URL

	apiHost string
	apiPath string
	apiKey  string

	stickersAPIPath string
	stickersAPIKey  string
}

// OptFunc is used to configure optional paramaters
type OptFunc func(*Client)

// New creates an API client with the specified options
func New(confs ...OptFunc) (*Client, error) {
	client := http.DefaultClient
	c := &Client{client: client}
	c.apiHost = apiHost
	c.apiPath = apiPath
	c.apiKey = apiKey

	c.stickersAPIPath = stickersAPIPath
	c.stickersAPIKey = apiKey

	for _, v := range confs {
		v(c)
	}

	u := fmt.Sprintf("https://%s/%s", c.apiHost, c.apiPath)
	ur, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	c.baseURL = ur

	su := fmt.Sprintf("https://%s/%s", c.apiHost, c.stickersAPIPath)
	sr, err := url.Parse(su)
	if err != nil {
		return nil, err
	}
	stick := stickerClient{client: c.client,
		baseURL: sr,
		apiHost: c.apiHost,
		apiPath: c.stickersAPIPath,
		apiKey:  c.stickersAPIKey,
	}
	c.Stickers = stick
	return c, nil
}

type StickerClient struct {
	client  *http.Client
	baseURL *url.URL

	apiHost string
	apiPath string
	apiKey  string
}
