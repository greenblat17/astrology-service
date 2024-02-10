package domain

type Star struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	ImageURL    string `json:"image_url"`
}
