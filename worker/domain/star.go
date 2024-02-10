package domain

type APOD struct {
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	HDURL       string `json:"hdurl"`
	MediaType   string `json:"media_type"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type Star struct {
	ID          int
	Name        string
	Description string
	Date        string
	ImageURL    string
}
