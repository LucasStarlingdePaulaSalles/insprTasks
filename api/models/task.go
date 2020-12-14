package models

import (
	"errors"
	"html"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Task is the applications's main class, the object that represenst a user task.
type Task struct {
	ID           uint64        `gorm:"primary_key;auto_increment" json:"id"`
	Title        string        `gorm:"size:255;not null;unique" json:"title"`
	Description  string        `gorm:"size:255;not null;" json:"description"`
	Priority     uint8         `gorm:"not null" json:"priority"`
	Status       uint8         `gorm:"not null" json:"status"`
	Deadline     time.Time     `gorm:"not null" json:"deadline"`
	TimeEstimate time.Duration `gorm:"not null" json:"timeEstimate"`
	Dependencies string        `gorm:"size:255;" json:"dependencies"`
	WorkedFrom   time.Time     `json:"workedFrom"`
	WorkedFor    time.Duration `json:"workedFor"`
	CreatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// Prepare is a default constructor like function that assigns default values for business and DB compatibility reasons
func (task *Task) Prepare() {
	task.ID = 0
	task.Title = html.EscapeString(strings.TrimSpace(task.Title))
	task.Description = html.EscapeString(strings.TrimSpace(task.Description))
	task.Status = 0
	task.Dependencies = html.EscapeString(strings.TrimSpace(task.Dependencies))
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
func (task *Task) WorkOnTask(db *gorm.DB, taskID uint64) (*Task, error) {

	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", taskID).Take(&task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("Incorrect task id")
		}
		return &Task{}, err
	}
	if task.Status == Working {
		return &Task{}, errors.New("Task already in progress")
	}

	err = task.verifyDependencies(db)
	if err != nil {
		return &Task{}, err
	}

	_, err = task.StopWorkOnTasks(db)
	if err != nil && err.Error() != "No task been worked on" {
		return &Task{}, err
	}

	err = nil
	db.Debug().Model(&Task{}).UpdateColumns(
		map[string]interface{}{
			"status":     1,
			"updatedAt":  time.Now(),
			"workedFrom": time.Now(),
		}).Take(&task)
	return task, err
}

func (task *Task) verifyDependencies(db *gorm.DB) error {
	dependencies := strings.Split(task.Dependencies, ";")
	var blockers []string
	for _, s := range dependencies {
		dependID, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return errors.New("Invalid dependencie Id")
		}
		err = db.Debug().Model(&Task{}).Where("id = ?", dependID).Where("status = ?", Done).Take(&Task{}).Error

		if err != nil {
			blockers = append(blockers, s)
		}
	}
	if len(blockers) > 0 {
		msg := "Task blocked. Blockers (by ID): "
		msg = msg + blockers[0]
		for _, s := range blockers[1:] {
			msg = msg + ", " + s
		}
		return errors.New(msg)
	}
	return nil
}

//StopWorkOnTasks interrupts work on the current task
func (task *Task) StopWorkOnTasks(db *gorm.DB) (*Task, error) {
	var err error
	var t = Task{}
	err = db.Debug().Model(&Task{}).Where("status = ?", 1).Take(&t).Error
	if err != nil {
		return &Task{}, errors.New("No task been worked on")
	}

	timeWorked := t.WorkedFor + time.Since(t.WorkedFrom)

	err = db.Debug().Model(&Task{}).Where("status = ?", 1).UpdateColumns(
		map[string]interface{}{
			"status":    0,
			"workedFor": timeWorked,
			"updatedAt": time.Now(),
		}).Error

	if err != nil {
		return &Task{}, err
	}
	db.Debug().Model(&Task{}).Take(&t)
	return &t, nil
}

//ChangeTaskStatus is a update function specifically for changin a task's status
func (task *Task) ChangeTaskStatus(db *gorm.DB, taskID uint64, newStatus uint8) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", taskID).Take(&task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("Incorrect task id")
		}
		return &Task{}, err
	}

	if task.Status == 1 {
		return task.StopWorkOnTasks(db)
	}

	if task.Status == 3 {
		return &Task{}, errors.New("Task closed")
	}

	err = db.Debug().Model(&Task{}).UpdateColumn("status", newStatus).Take(&task).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("Incorrect task id")
		}
		return &Task{}, err
	}

	return task, err
}

func dateCompare(task Task, filter DateFilterDTI) bool {
	var date time.Time
	switch filter.Field {
	case "deadline":
		date = task.Deadline
	default:
		date = task.CreatedAt
	}
	return date.Year() == filter.Year && int(date.Month()) == filter.Month && date.Day() == filter.Day
}

//FindByDate return all tasks that match the given DateFilter
func (task *Task) FindByDate(db *gorm.DB, filter DateFilterDTI) (*[]Task, error) {
	var err error
	tasks := []Task{}
	err = db.Debug().Model(&Task{}).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	response := []Task{}
	for i := range tasks {
		if dateCompare(tasks[i], filter) {
			response = append(response, tasks[i])
		}
	}
	return &response, err
}

func valueCompare(task Task, filter NumericFilterDTI) bool {
	var value uint8
	switch filter.Field {
	default:
		value = task.Priority
	}
	return value == filter.Value
}

//FindByValue return all tasks that match the given DateFilter
func (task *Task) FindByValue(db *gorm.DB, filter NumericFilterDTI) (*[]Task, error) {
	var err error
	tasks := []Task{}
	err = db.Debug().Model(&Task{}).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	response := []Task{}
	for i := range tasks {
		if valueCompare(tasks[i], filter) {
			response = append(response, tasks[i])
		}
	}
	return &response, err
}