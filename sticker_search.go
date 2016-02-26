package gifgo

import (
	"encoding/json"
	"net/http"
)

// Search searches for a sticker GIF with the specified params
func (c *stickerClient) Search(query SearchReq) (*MultipleGIF, error) {
	reqURL := c.baseURL.ResolveReference(searchPath)
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
