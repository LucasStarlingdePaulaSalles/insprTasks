package models

//NumericFilterDTI is a Data Type Input for filtering tasks by a numeric field
type NumericFilterDTI struct {
	Field string `json:"field"`
	Value int    `json:"value"`
}
