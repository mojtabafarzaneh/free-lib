package models

import "net/url"

type SearchQuery struct {
	Query        string `json:"query"`
	SearchMirror url.URL
	Print        bool
}

type DownloadQuery struct {
	Hash string `json:"hash"`
}
