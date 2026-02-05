package notes

import "time"

const ROUTES_FOLDER = "src/routes/(notes)/"

type Note struct {
	Title string    `json:"title"`
	Slug  string    `json:"slug"`
	Date  time.Time `json:"dates"`
}

type FrontMatter struct {
	Title string    `yaml:"title"`
	Date  time.Time `yaml:"date"`
}
