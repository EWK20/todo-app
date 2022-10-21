package dto

type SearchReq struct {
	Item      string `json:"item" filter:"searchable"`
	Completed bool   `json:"completed" filter:"searchable"`
}
