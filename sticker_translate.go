package gifgo

import (
	"encoding/json"
	"net/http"
)

// Translate runs giphy's sticker "translate" on a term. ~magic~
func (c *StickerClient) Translate(query TranslateReq) (*SingleGIF, error) {
	reqURL := c.baseURL.ResolveReference(translatePath)
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
	data := new(SingleGIF)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
