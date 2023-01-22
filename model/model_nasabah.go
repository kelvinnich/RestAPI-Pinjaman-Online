package model

type Nasabah struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Nama string `gorm:"type:varchar(255)" json:"nama"`
	Email string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	NoTelpon string `gorm:"type:varchar(255)" json:"no_telpon"`
	Alamat string `gorm:"type:varchar(255)" json:"alamat"`
	NoKtp string `gorm:"uniqueIndex;type:varchar(255)" json:"no_ktp"`
	StatusVerified bool `gorm:"not null;default:false" json:"status_verified"`
	Pekerjaan Pekerjaan_nasabah `gorm:"foreignkey:NasabahID"`
}
