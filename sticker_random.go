package gifgo

import (
	"encoding/json"
	"net/http"
)

// Random returns a random sticker GIF
func (c *stickerClient) Random(query RandomReq) (*RandomGIF, error) {
	reqURL := c.baseURL.ResolveReference(randomPath)
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

	data := new(RandomGIF)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
