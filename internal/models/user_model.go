package models

import "time"

type UserModel struct {
	Id         int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Username   string    `gorm:"column:username;size:255; null" json:"username"`
	Password   string    `gorm:"column:password;size:255;null" json:"password"`
	Email      string    `gorm:"column:email;size:255;null" json:"email"`
	CreatedAt  time.Time `gorm:"column:email;size:255;null" json:"created_at"`
	UpdateddAt time.Time `gorm:"column:email;size:255;null" json:"updated_at"`
}

func (UserModel) TableName() string {
	return "user"
}
