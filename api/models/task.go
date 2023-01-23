package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint      `gorm:"primary_key;auto_increment"json:"id"`
	UserID      uint      `gorm:"not null"json:"user_id"`
	Title       string    `gorm:"type:varchar(100);not null"json:"title"`
	Description string    `gorm:"type:varchar(100);not null"json:"description"`
	StartDate   time.Time `gorm:"type:varchar(100);not null"json:"start_date"`
	EndDate     time.Time `gorm:"type:varchar(100);not null"json:"end_date"`
}

func (t *Task) Prepare() {
	t.ID = 0
	t.UserID = t.UserID
	t.Title = html.EscapeString(strings.TrimSpace(t.Title))
	t.Description = html.EscapeString(strings.TrimSpace(t.Description))
	t.StartDate = time.Now()
	t.EndDate = time.Now()
}

func (t *Task) ValidateTask(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if t.Title == "" {
			return errors.New("Required Title")
		}
		if t.Description == "" {
			return errors.New("Required Description")
		}
		if string(t.UserID) == "" {
			return errors.New("Required UserID")
		}
		return nil
	default:
		if t.Title == "" {
			return errors.New("Required Title")
		}
		if t.Description == "" {
			return errors.New("Required Description")
		}
		if string(t.UserID) == "" {
			return errors.New("Required UserID")
		}
		return nil
	}
}

func (t *Task) SaveTask(db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Create(&t).Error
	if err != nil {
		return &Task{}, err
	}
	return t, nil
}
func (t *Task) FindTaskByUserID(db *gorm.DB, uid uint) (*[]Task, error) {
	var err error
	tasks := []Task{}
	err = db.Debug().Model(&Task{}).Where("user_id = ?", uid).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	return &tasks, nil
}
func (t *Task) FindTaskByID(db *gorm.DB, tid uint) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", tid).Take(&t).Error
	if err != nil {
		return &Task{}, err
	}
	return t, nil
}
func (t *Task) UpdateATask(db *gorm.DB, tid uint) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", tid).Updates(Task{Title: t.Title, Description: t.Description, StartDate: t.StartDate, EndDate: t.EndDate}).Error
	if err != nil {
		return &Task{}, err
	}
	return t, nil
}
func (t *Task) DeleteATask(db *gorm.DB, tid uint) (int64, error) {
	db = db.Debug().Model(&Task{}).Where("id = ?", tid).Take(&Task{}).Delete(&Task{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
