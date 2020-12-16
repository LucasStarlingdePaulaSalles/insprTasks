package models

//DateFilterDTO is a Data Type Input for filtering tasks by date
type DateFilterDTO struct {
	Field string `json:"field"`
	Year  int    `json:"year"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}
