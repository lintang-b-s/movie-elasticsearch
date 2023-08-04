package entity

type Movie struct {
	ReleaseYear string `json:"releaseYear"`
	Title       string `json:"title"`
	Director    string `json:"director"`
	Cast        string `json:"cast"`
	Genre       string `json:"genre"`
	Synopsis    string `json:"synopsis"`
}
