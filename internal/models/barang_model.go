package models

import "time"

type BarangModel struct {
	Id         int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Nama       string    `gorm:"column:nama;size:255;not null" json:"nama"`
	Kode       string    `gorm:"column:kode;size:100" json:"kode"`
	KategoryId int       `gorm:"column:kategory_id" json:"kategory_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (BarangModel) TableName() string {
	return "barang"
}
