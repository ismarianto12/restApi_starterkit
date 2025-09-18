package models

import "time"

type PurcashingModel struct {
	Id         int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Nama       string    `gorm:"column:nama;size:50;not null" json:"nama"`
	Harga      string    `gorm:"column:harga;size:50;not null" json:"harga"`
	ApproveId  string    `gorm:"column:approved_id;size:50;not null" json:"approved_id"`
	Kode       string    `gorm:"column:kode;size:100" json:"kode"`
	KategoryId int       `gorm:"column:kategory_id" json:"kategory_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

// Nama tabel sesuai dengan yang ada di DB
func (PurcashingModel) TableName() string {
	return "purchasing_table"
}
