package gifgo

// Random returns a random sticker GIF
func (c *StickerClient) Random(query RandomReq) (*RandomGIF, error) {
	reqURL := c.baseURL.ResolveReference(randomPath)
	return random(query, reqURL, c.apiKey, c.client)
}
