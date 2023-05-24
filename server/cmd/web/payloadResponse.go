package web

type MoviePayloadResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Year        string   `json:"year"`
	ReleaseDate string   `json:"release_date"`
	Runtime     string   `json:"runtime"`
	Rating      string   `json:"rating"`
	MPAARating  string   `json:"mpaa_rating"`
	GenreID     []string `json:"genre_id"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}