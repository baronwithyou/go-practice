package poster

type Request struct {
	Title string `json:"s"`
}

type MovieInfo struct {
	Title      string
	Year       string
	Actor      string
	Poster     string
	ImdbRating string `json:"imdbRating"`
}
