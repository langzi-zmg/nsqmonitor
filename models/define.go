package models


type Pagination struct {
	Page  int64 `json:"page" query:"page"`
	Limit int64 `json:"limit" query:"limit"`
}

