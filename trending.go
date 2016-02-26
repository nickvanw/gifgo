package gifgo

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Trending returns the current trending GIFs
func (c *Client) Trending(query TrendingReq) (*MultipleGIF, error) {
	reqURL := c.baseURL.ResolveReference(trendingPath)
	q := query.toValues()
	q.Add("api_key", c.apiKey)
	reqURL.RawQuery = q.Encode()

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
