package model

import "time"

type Pembayaran struct {
	ID int` gorm:"primary_key;column:id;type:serial" json:"id"`
	Pinjaman_id int `gorm:"not null" json:"-"`
	Pinjaman Pinjaman `gorm:"foreignKey:Pinjaman_id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"pinjaman_id"`
	Jumlah int `gorm:"type:integer" json:"jumlah"`
	Pembayaran_perbulan int `gorm:"type:integer" json:"pembayaran_perbulan"`
	Status_Pembayaran bool `gorm:"type:boolean" json:"status_pembayaran"`
	Tanggal_Pembayaran time.Time `gorm:"type:timestamp" json:"tanggal_pembayaran"`
	}