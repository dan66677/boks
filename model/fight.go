package model

type Fight struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Fighter1 string `json:"fighter1"`
	Fighter2 string `json:"fighter2"`
	Winner   string `json:"winner"`
}
