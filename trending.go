package gifgo

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Trending returns the current trending GIFs
func (c *Client) Trending(query TrendingReq) (*MultipleGIF, error) {
	reqURL := c.baseURL.ResolveReference(trendingPath)
	return trending(query, reqURL, c.apiKey, c.client)
}

// TrendingReq contains info for a trending GIF search
type TrendingReq struct {
	Limit  int
	Rating Rating
}

func (s TrendingReq) toValues() url.Values {
	values := url.Values{}
	if s.Limit != 0 {
		values.Add("limit", string(s.Limit))
	}
	if s.Rating != "" {
		values.Add("rating", string(s.Rating))
	}
	return values
}

func trending(query TrendingReq, reqURL *url.URL, key string, client *http.Client) (*MultipleGIF, error) {
	q := query.toValues()
	q.Add("api_key", key)
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data := new(MultipleGIF)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
