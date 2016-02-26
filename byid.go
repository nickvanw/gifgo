package gifgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GIFByID returns a single GIF by it's giphy ID
func (c *Client) GIFByID(id string) (*SingleGIF, error) {
	path, err := url.Parse(id)
	if err != nil {
		return nil, err
	}
	reqURL := c.baseURL.ResolveReference(path)

	q := reqURL.Query()
	q.Add("api_key", c.apiKey)
	reqURL.RawQuery = q.Encode()

	fmt.Println(reqURL.String())

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	data := new(SingleGIF)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

// GIFsByID returns 1 or more GIFs based on the IDs passed in
func (c *Client) GIFsByID(ids ...string) (*MultipleGIF, error) {
	reqURL := c.baseURL.ResolveReference(idsPath)
	q := reqURL.Query()
	q.Add("api_key", c.apiKey)
	idstr := strings.Join(ids, ",")
	q.Add("ids", idstr)
	reqURL.RawQuery = q.Encode()

	fmt.Println(reqURL.String())

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	data := new(MultipleGIF)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
