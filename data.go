package gifgo

// Rating represents filter params for filtering by rating
type Rating string

// These constants represent the different ratings
const (
	YRated    Rating = "y"
	GRated    Rating = "g"
	PGRated   Rating = "pg"
	PG13Rated Rating = "pg-13"
	RRated    Rating = "r"
)

// SingleGIF represents single GIF return data
type SingleGIF struct {
	Data GIF  `json:"data"`
	Meta Meta `json:"meta"`
}

// MultipleGIF represents one or more returned GIFs
type MultipleGIF struct {
	Data       []GIF `json:"data"`
	Meta       Meta  `json:"meta"`
	Pagination struct {
		TotalCount int `json:"total_count"`
		Count      int `json:"count"`
		Offset     int `json:"offset"`
	} `json:"pagination"`
}

// RandomGIF is a random GIF
type RandomGIF struct {
	Data struct {
		Type                         string `json:"type"`
		ID                           string `json:"id"`
		URL                          string `json:"url"`
		ImageOriginalURL             string `json:"image_original_url"`
		ImageURL                     string `json:"image_url"`
		ImageMp4URL                  string `json:"image_mp4_url"`
		ImageFrames                  string `json:"image_frames"`
		ImageWidth                   string `json:"image_width"`
		ImageHeight                  string `json:"image_height"`
		FixedHeightDownsampledURL    string `json:"fixed_height_downsampled_url"`
		FixedHeightDownsampledWidth  string `json:"fixed_height_downsampled_width"`
		FixedHeightDownsampledHeight string `json:"fixed_height_downsampled_height"`
		FixedWidthDownsampledURL     string `json:"fixed_width_downsampled_url"`
		FixedWidthDownsampledWidth   string `json:"fixed_width_downsampled_width"`
		FixedWidthDownsampledHeight  string `json:"fixed_width_downsampled_height"`
		FixedHeightSmallURL          string `json:"fixed_height_small_url"`
		FixedHeightSmallStillURL     string `json:"fixed_height_small_still_url"`
		FixedHeightSmallWidth        string `json:"fixed_height_small_width"`
		FixedHeightSmallHeight       string `json:"fixed_height_small_height"`
		FixedWidthSmallURL           string `json:"fixed_width_small_url"`
		FixedWidthSmallStillURL      string `json:"fixed_width_small_still_url"`
		FixedWidthSmallWidth         string `json:"fixed_width_small_width"`
		FixedWidthSmallHeight        string `json:"fixed_width_small_height"`
		Username                     string `json:"username"`
		Caption                      string `json:"caption"`
	} `json:"data"`
	Meta Meta `json:"meta"`
}

// GIF is a single GIF data, with one or more images
type GIF struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	Slug             string `json:"slug"`
	URL              string `json:"url"`
	BitlyGIFURL      string `json:"bitly_gif_url"`
	BitlyURL         string `json:"bitly_url"`
	EmbedURL         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Rating           string `json:"rating"`
	ContentURL       string `json:"content_url"`
	SourceTLD        string `json:"source_tld"`
	SourcePostURL    string `json:"source_post_url"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`

	User   User   `json:"user"`
	Images Images `json:"images"`
}

// User represents a user (from where?)
type User struct {
	AvatarURL   string `json:"avatar_url"`
	BannerURL   string `json:"banner_url"`
	ProfileURL  string `json:"profile_url"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Twitter     string `json:"twitter"`
}

// Images is a list of different representations of the same images saved differently
type Images struct {
	Original               Image `json:"original"`
	FixedHeight            Image `json:"fixed_height"`
	FixedHeightStill       Image `json:"fixed_height_still"`
	FixedHeightDownsampled Image `json:"fixed_height_downsampled"`
	FixedWidth             Image `json:"fixed_width"`
	FixedWidthStill        Image `json:"fixed_width_still"`
	FixedWidthDownsampled  Image `json:"fixed_width_downsampled"`
	FixedHeightSmall       Image `json:"fixed_height_small"`
	FixedHeightSmallStill  Image `json:"fixed_height_small_still"`
	FixedWidthSmall        Image `json:"fixed_width_small"`
	FixedWidthSmallStill   Image `json:"fixed_width_small_still"`
	Downsized              Image `json:"downsized"`
	DownsizedStill         Image `json:"downsized_still"`
	DownsizedLarge         Image `json:"downsized_large"`
	DownsizedMedium        Image `json:"downsized_medium"`
	OriginalStill          Image `json:"original_still"`
	Looping                Image `json:"looping"`
}

// Image represents a single image
type Image struct {
	URL      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Frames   string `json:"frames"`
	MP4      string `json:"mp4"`
	MP4Size  string `json:"mp4_size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}

// Meta contains information about the response
type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
