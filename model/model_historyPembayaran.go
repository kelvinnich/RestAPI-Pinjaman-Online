package model

import "time"

type HistoryPembayaran struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Pinjaman_id int `gorm:"not null" json:"-"`
	IDPembayaran int `gorm:"primary_key"json:"-"`
	Tanggal time.Time `gorm:"not null"`
	Pinjaman Pinjaman `gorm:"foreignKey:Pinjaman_id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"pinjaman_id"`
	Pembayaran Pembayaran `gorm:"foreignKey:IDPembayaran;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"pembayaran_id"`
}