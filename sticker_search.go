package gifgo

// Search searches for a sticker GIF with the specified params
func (c *StickerClient) Search(query SearchReq) (*MultipleGIF, error) {
	reqURL := c.baseURL.ResolveReference(searchPath)
	return search(query, reqURL, c.apiKey, c.client)
}
