package gifgo

// Trending returns the current trending sticker GIFs
func (c *StickerClient) Trending(query TrendingReq) (*MultipleGIF, error) {
	reqURL := c.baseURL.ResolveReference(trendingPath)
	return trending(query, reqURL, c.apiKey, c.client)
}
