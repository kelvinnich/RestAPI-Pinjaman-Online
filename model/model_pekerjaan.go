package model

import "time"

type Pekerjaan_nasabah struct {
	ID         int `gorm:"primary_key;auto_increment" json:"id"`
	NasabahID  int ` gorm:"not null " json:"-"`
	AlamatPerusahaan string `gorm:"type:varchar(255)" json:"alamat_perusahaan"`
	TanggalGajian time.Time `gorm:"type:timestamp" json:"tanggal_gajian"`
	PosisiPekerjaan string `gorm:"type:varchar(255)" json:"posisi_perkerjaan"`
}