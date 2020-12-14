package models

//DateFilter is a filter structure for date comparissons
type DateFilter struct {
	Field string `json:"field"`
	Year int     `json:"year"`
	Month int    `json:"month"`
	Day int      `json:"day"`
}