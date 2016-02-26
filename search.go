package gifgo

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Search searches for a GIF with the specified params
func (c *Client) Search(query SearchReq) (*MultipleGIF, error) {
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

// SearchReq contains paramaters for a GIF search
type SearchReq struct {
	Query  string
	Limit  int
	Offset int
	Rating string
}

func (s SearchReq) toValues() url.Values {
	values := url.Values{}
	values.Add("q", s.Query)
	if s.Limit != 0 {
		values.Add("limit", string(s.Limit))
	}
	if s.Offset != 0 {
		values.Add("offset", string(s.Offset))
	}
	if s.Rating != "" {
		values.Add("rating", s.Rating)
	}
	return values
}
