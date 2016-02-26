package gifgo

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Random returns a random GIF
func (c *Client) Random(query RandomReq) (*RandomGIF, error) {
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

// RandomReq contains options for listing a random GIF
type RandomReq struct {
	Tag    string
	Rating Rating
}

func (s RandomReq) toValues() url.Values {
	values := url.Values{}
	values.Add("tag", string(s.Tag))

	if s.Rating != "" {
		values.Add("rating", string(s.Rating))
	}
	return values
}
