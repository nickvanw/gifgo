package gifgo

import (
	"encoding/json"
	"net/http"
)

// Trending returns the current trending sticker GIFs
func (c *stickerClient) Trending(query TrendingReq) (*MultipleGIF, error) {
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
