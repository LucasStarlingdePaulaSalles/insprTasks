package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Task is the applications's main class, the object that represenst a user task.
type Task struct {
	ID           uint64        `gorm:"primary_key;auto_increment" json:"id"`
	Title     	 string        `gorm:"size:255;not null;unique" json:"title"`
	Description  string        `gorm:"size:255;not null;" json:"description"`
	Priority     uint16        `gorm:"not null" json:"priority"`
	Status       uint8         `gorm:"not null" json:"status"`
	Deadline     time.Time     `gorm:"not null" json:"deadline"`
	TimeEstimate time.Duration `gorm:"not null" json:"timeEstimate"`
	WorkedFrom   time.Time     `json:"workedFrom"`
	WorkedFor    time.Duration `json:"workedFor"`
	CreatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// Prepare is a default constructor like function
func (task *Task) Prepare() {
	task.ID = 0
	task.Title = html.EscapeString(strings.TrimSpace(task.Title))
	task.Description = html.EscapeString(strings.TrimSpace(task.Description))
	task.Status = 0
	task.WorkedFor = 0
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
}

func (task *Task) Validate() error {

	if task.Title == "" {
		return errors.New("Title is required")
	}
	if task.Description == "" {
		return errors.New("Description is required")
	}
	if task.TimeEstimate < 1 {
		return errors.New("Time Estimate is required")
	}
	if task.Priority < 1 {
		return errors.New("A task must have some level of priority (bigger than 0)")
	}
	return nil
}

func (task *Task) SaveTask(db *gorm.DB) (*Task, error) {

	var err error
	err = db.Debug().Create(&task).Error
	if err != nil {
		return &Task{}, err
	}
	return task, nil
}

func (task *Task) FindAllTasks(db *gorm.DB) (*[]Task, error) {
	var err error
	tasks := []Task{}
	err = db.Debug().Model(&Task{}).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	return &tasks, err
}