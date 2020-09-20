package models

type Pagination struct {
	Page uint `json:"page"`
	Per uint `json:"per"`
	Count uint `json:"count"`
}

func ToOffset(pagination *Pagination) uint {
	return pagination.Page - 1
}

