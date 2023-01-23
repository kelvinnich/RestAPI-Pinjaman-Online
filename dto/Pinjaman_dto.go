package dto

type CreatePinjamanDTO struct {
	Nasabah_id uint64 `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	Jumlah int `json:"jumlah" form:"jumlah" binding:"required"`
	SukuBunga int `json:"suku_bunga" form:"suku_bunga"binding:"required"`
	Durasi int `json:"durasi_peminjaman" form:"durasi" binding:"required"`
}

type UpdatePinjamanDTO struct {
	Id uint64 `json:"id" form:"id"`
	Nasabah_id uint64 `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	Jumlah int `json:"jumlah" form:"jumlah" binding:"required"`
	SukuBunga int `json:"suku_bunga" form:"suku_bunga"binding:"required"`
	Durasi int `json:"durasi_peminjaman" form:"durasi" binding:"required"`
}
