package gifgo

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Translate runs giphy's "translate" on a term. ~magic~
func (c *Client) Translate(query TranslateReq) (*SingleGIF, error) {
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

// TranslateReq contains the paramaters for a translate request
type TranslateReq struct {
	Term   string
	Rating Rating
}

func (s TranslateReq) toValues() url.Values {
	values := url.Values{}
	values.Add("s", s.Term)
	if s.Rating != "" {
		values.Add("rating", string(s.Rating))
	}
	return values
}
