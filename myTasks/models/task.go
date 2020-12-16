package models

import "time"

//Task is class for representing tasks.
type Task struct {
	ID           uint64        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Priority     uint8         `json:"priority"`
	Status       uint8         `json:"status"`
	Deadline     time.Time     `json:"deadline"`
	TimeEstimate time.Duration `json:"timeEstimate"`
	Dependencies string        `json:"dependencies"`
	WorkedFrom   time.Time     `json:"workedFrom"`
	WorkedFor    time.Duration `json:"workedFor"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
}
