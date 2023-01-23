package model

type Pinjaman struct {
	Id uint64 `gorm:"primary_key;auto_increment;column:id" json:"id"`
	IDNasabah uint64 `gorm:"not null" json:"-"`
	Nasabah Nasabah `gorm:"foreignKey:IDNasabah;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"nasabah_id"`
	Jumlah int `gorm:"type:integer" json:"jumlah"`
	SukuBunga int `gorm:"type:integer" json:"suku_bunga"`
	Durasi int `gorm:"type:integer" json:"durasi_peminjaman"`
	StatusApproved bool `gorm:"type:boolean" json:"status_approved"`
}
