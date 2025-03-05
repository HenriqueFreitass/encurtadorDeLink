package models

type Shortener struct {
	Id          string
	SiteName    string
	OriginalUrl string
	NewUrl      string
	Views       int
	UserEmail   string
}
