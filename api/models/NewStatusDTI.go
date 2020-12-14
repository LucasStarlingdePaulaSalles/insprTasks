package models

//NewStatusDTI is a Data Type Input for updating a task's status
type NewStatusDTI struct{
	NewStatus uint8 `json:newStatus`
}