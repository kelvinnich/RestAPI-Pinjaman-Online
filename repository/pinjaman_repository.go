package repository

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type PinjamanRepository interface{
	CreatePinjamanRepository(pinjaman *model.Pinjaman) error
	UpdatePinjamanRepository(id int, pinjaman *model.Pinjaman)error
	SearchPinjamanByIdRepository(id int) (*model.Pinjaman, error)
	DeletePinjamanRepository(id int)error
	GetTotalPembayaranNasabah(nasabahID int) (int, error)
}

type pinjamanConnection struct {
	db *gorm.DB
}

func NewPinjamanRepository(db *gorm.DB)PinjamanRepository{
	return &pinjamanConnection{
		db: db,
	}
}

func(db *pinjamanConnection)CreatePinjamanRepository(pinjaman *model.Pinjaman)error{
	if err := db.db.Create(pinjaman).Error; err != nil {
		return err
	}
	return nil
}

func(db *pinjamanConnection)UpdatePinjamanRepository(id int, pinjaman *model.Pinjaman)error{
	if err := db.db.Model(model.Pinjaman{}).Where("id = $1").Updates(pinjaman).Error; err != nil {
		return err
	}
	return nil
}

func(db *pinjamanConnection)SearchPinjamanByIdRepository(id int) (*model.Pinjaman, error){
	var pinjaman model.Pinjaman
	if err := db.db.First(&pinjaman, id).Error; err != nil{
		return nil,err
	}

	return &pinjaman,nil
}

func(db *pinjamanConnection)DeletePinjamanRepository(id int)error{
	if err := db.db.Where("id = $1", id).Delete(&model.Pinjaman{}).Error; err != nil{
		return err
	}

	return nil
}

func (db *pinjamanConnection) GetTotalPembayaranNasabah(nasabahID int) (int, error) {
	var pinjaman model.Pinjaman
	if err := db.db.Where("nasabah_id = $1", nasabahID).Find(&pinjaman).Error; err != nil {
			return 0, err
	}
	jumlahPinjaman := pinjaman.Jumlah
	sukuBunga := pinjaman.SukuBunga
	durasiPinjaman := pinjaman.Durasi
	pembayaranPerBulan := (jumlahPinjaman * sukuBunga) / (12 * 100) + (jumlahPinjaman / durasiPinjaman)
	totalPembayaran := pembayaranPerBulan * durasiPinjaman
	return totalPembayaran, nil
}
