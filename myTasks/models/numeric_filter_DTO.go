package models

//NumericFilterDTO is a Data Type Input for filtering tasks by a numeric field
type NumericFilterDTO struct {
	Field string `json:"field"`
	Value int    `json:"value"`
}
