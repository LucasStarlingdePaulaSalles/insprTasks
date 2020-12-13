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

var NoCurrentWorkError = errors.New("No task been worked on")
// Prepare is a default constructor like function that assigns default values for business and DB compatibility reasons
func (task *Task) Prepare() {
	task.ID = 0
	task.Title = html.EscapeString(strings.TrimSpace(task.Title))
	task.Description = html.EscapeString(strings.TrimSpace(task.Description))
	task.Status = 0
	task.WorkedFor = 0 * time.Nanosecond
	task.WorkedFrom = time.Now()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
}

//Validate is a Validator intended to inforce business rules on creating or updating Tasks
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

//SaveTask is responsible for storing a task on the DB
func (task *Task) SaveTask(db *gorm.DB) (*Task, error) {

	var err error
	err = db.Debug().Create(&task).Error
	if err != nil {
		return &Task{}, err
	}
	return task, nil
}

//FindAllTasks queries the DB to select al availible Tasks
func (task *Task) FindAllTasks(db *gorm.DB) (*[]Task, error) {
	var err error
	tasks := []Task{}
	err = db.Debug().Model(&Task{}).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	return &tasks, err
}

//WorkOnTask iniciates time counting for a specified task.
func (task *Task) WorkOnTask(db *gorm.DB, taskID uint32) (*Task, error){

	var err error
	err = db.Debug().Model(Task{}).Where("id = ?",taskID).Take(&task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("Incorrect task id")
		}
		return &Task{}, err
	}
	if task.Status == 1 {
		return &Task{}, errors.New("Task already in progress")
	}
	_,err = task.StopWorkOnTasks(db)
	if err != nil && err.Error() != "No task been worked on" {
		return &Task{}, err
	}
	err = nil
	db.Debug().Model(Task{}).Where("id = ?",taskID).UpdateColumns(
		map[string]interface{}{
			"status": 1,
			"updatedAt": time.Now(),
			"workedFrom": time.Now(),
			}).Take(&task)
	return task, err
}

//StopWorkOnTasks interrupts work on the current task
func (task *Task) StopWorkOnTasks(db *gorm.DB) (*Task, error){
	var err error
	var t = Task{}
	err = db.Debug().Model(&Task{}).Where("status = ?", 1).Take(&t).Error
	if err != nil {
		return &Task{}, errors.New("No task been worked on")
	}

	timeWorked := t.WorkedFor + time.Since(t.WorkedFrom)

	err = db.Debug().Model(Task{}).Where("status = ?", 1).UpdateColumns(
		map[string]interface{}{
			"status": 0,
			"workedFor": timeWorked,
			"updatedAt": time.Now(),
		}).Error

	if err != nil {
		return &Task{}, err
	}
	db.Debug().Model(Task{}).Take(&t)
	return &t, nil
}