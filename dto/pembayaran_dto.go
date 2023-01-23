package dto

import "time"

type CreatePembayaranDTO struct {
	PinjamanID int `form:"pinjaman_id" json:"pinjaman_id" binding:"required"`
	Jumlah int `form:"jumlah" json:"jumlah" binding:"required"`
	PembayaranPerbulan int `form:"pembayaran_perbulan" json:"pembayaran_perbulan" binding:"required"`
	TanggalPembayaran time.Time `form:"tanggal_pembayaran" json:"tanggal_pembayaran" binding:"required"`
}

type UpdatePembayaranDTO struct {
	Id int `form:"id" json:"id"`
	Jumlah int `form:"jumlah" json:"jumlah"`
	PembayaranPerbulan int `form:"pembayaran_perbulan" json:"pembayaran_perbulan"`
	TanggalPembayaran time.Time `form:"tanggal_pembayaran" json:"tanggal_pembayaran"`
	StatusPembayaran bool `form:"status_pembayaran" json:"status_pembayaran"`
}

