package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id             int       `json:"id" gorm:"primaryKey"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"created_at"`
	LastLoggedInAt time.Time `json:"last_logged_in_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = u.Id
	response["email"] = u.Email
	response["created_at"] = u.CreatedAt

	return response
}
