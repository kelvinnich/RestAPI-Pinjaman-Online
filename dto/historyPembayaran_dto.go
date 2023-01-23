package dto

import "time"


type CreateHistoryPembayaranDTO struct {
	Pinjaman_id int `json:"pinjaman_id" form:"pinjaman_id" binding:"required"`
	IDPembayaran int `json:"pembayaran_id" form:"pembayaran_id" binding:"required"`
	Tanggal time.Time `json:"tanggal" form:"tanggal" binding:"required"`
}

type UpdateHistoryPembayaranDTO struct {
	Id uint64 `json:"id" form:"id"`
	Pinjaman_id int `json:"pinjaman_id" form:"pinjaman_id" binding:"required"`
	IDPembayaran int `json:"pembayaran_id" form:"pembayaran_id" binding:"required"`
	Tanggal time.Time `json:"tanggal" form:"tanggal" binding:"required"`
}
