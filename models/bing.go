package models

type SearchResult struct {
	Keyword string `storm:"index"`
	URL     string `storm:"id"`
}
