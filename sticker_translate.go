package gifgo

// Translate runs giphy's sticker "translate" on a term. ~magic~
func (c *StickerClient) Translate(query TranslateReq) (*SingleGIF, error) {
	reqURL := c.baseURL.ResolveReference(translatePath)
	return translate(query, reqURL, c.apiKey, c.client)
}
