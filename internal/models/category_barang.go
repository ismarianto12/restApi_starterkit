package models

import "time"

type CategoryBarang struct {
	Id         int       `gorm:"primaryKey;autoIncrement;column:id"`
	Nama       string    `gorm:"column:nama;size:255;not null"`
	Kode       string    `gorm:"column:kode;size:100"`
	KategoryId int       `gorm:"column:kategory_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (CategoryBarang) TableName() string {
	return "category_barang"
}
