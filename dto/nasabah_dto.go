package dto

type RegisterNasabahDTO struct{
	Nama string `json:"nama" form:"nama" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	NoTelpon string `json:"no_telpon" form:"no_telpon" binding:"required"`
	Alamat string `json:"alamat" form:"alamat" binding:"required"`
	NoKtp string `json:"no_ktp" form:"no_ktp" binding:"required"`
}

type LoginNasabahDTO struct{
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateNasabahDTO struct{
	Id uint64 `json:"id" form:"id"`
	Nama string `json:"nama" form:"nama" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	NoTelpon string `json:"no_telpon" form:"no_telpon" binding:"required"`
	Alamat string `json:"alamat" form:"alamat" binding:"required"`
	NoKtp string `json:"no_ktp" form:"no_ktp" binding:"required"`
}