package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary key:auto_increment"json:"id"`
	Username  string    `gorm:"type:varchar(100);not null;unique"json:"username"`
	Email     string    `gorm:"type:varchar(100);not null;unique"json:"email"`
	Password  string    `gorm:"type:varchar(100);not null"json:"password"`
	CreatedAt time.Time `gorm:"default:'0000-00-00 00:00:00'"json:"created_at"`
	UpdatedAt time.Time `gorm:"default:'0000-00-00 00:00:00'"json:"updated_at"`
}

var err error

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil
	default:
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}
}
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.ErrRecordNotFound == err {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Updates(User{Username: u.Username, Email: u.Email, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &User{}, err
	}
	return u, err
}
func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		if gorm.ErrRecordNotFound == db.Error {
			return 0, errors.New("User Not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
