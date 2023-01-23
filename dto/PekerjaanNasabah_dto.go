package dto

import "time"


type CreatePekerjaanNasabahDTO struct{
	NasabahId int `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	AlamatPerusahaan string `json:"alamat_perusahaan" form:"alamat_perusahaan" binding:"required"`
	TanggalGajian time.Time `json:"tanggal_gajian" form:"tanggal_gajian" binding:"required"`
	PosisiPekerjaan string `json:"posisi_pekerjaan" form:"posisi_pekerjaan" binding:"required"`
}

type UpdatePekerjaanNasabahDTO struct{
	Id int `json:"id" form:"id"`
	NasabahId int `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	AlamatPerusahaan string `json:"alamat_perusahaan" form:"alamat_perusahaan" binding:"required"`
	TanggalGajian time.Time `json:"tanggal_gajian" form:"tanggal_gajian" binding:"required"`
	PosisiPekerjaan string `json:"posisi_pekerjaan" form:"posisi_pekerjaan" binding:"required"`
}