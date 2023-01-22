package model

type Dokumen_nasabah struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	IdNasabah uint `gorm:"not null " json:"-"`
	Nasabah_id Nasabah `gorm:"foreignKey:IdNasabah;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"nasabah_id"`
	JenisDokumen string `gorm:"type:varchar(255)" json:"jenis_dokumen"`
	FilePath string `gorm:"type:varchar(255)" json:"file_path"`
}
