package repository

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type PinjamanRepository interface{
	CreatePinjamanRepository(pinjaman *model.Master_Loan) error
	UpdatePinjamanRepository(id uint64, pinjaman *model.Master_Loan)error
	SearchPinjamanByIdRepository(id int) (*model.Master_Loan, error)
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

func(db *pinjamanConnection)CreatePinjamanRepository(pinjaman *model.Master_Loan)error{
	if err := db.db.Create(pinjaman).Error; err != nil {
		return err
	}
	return nil
}

func(db *pinjamanConnection)UpdatePinjamanRepository(id uint64, pinjaman *model.Master_Loan)error{
	if err := db.db.Model(model.Master_Loan{Id: id}).Updates(pinjaman).Error; err != nil {
		return err
	}
	return nil
}

func(db *pinjamanConnection)SearchPinjamanByIdRepository(id int) (*model.Master_Loan, error){
	var pinjaman model.Master_Loan
	if err := db.db.First(&pinjaman, id).Error; err != nil{
		return nil,err
	}

	return &pinjaman,nil
}

func(db *pinjamanConnection)DeletePinjamanRepository(id int)error{
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Loan{}).Error; err != nil{
		return err
	}

	return nil
}

func (db *pinjamanConnection) GetTotalPembayaranNasabah(nasabahID int) (int, error) {
	var pinjaman model.Master_Loan
	if err := db.db.Where("customer_id = $1", nasabahID).Find(&pinjaman).Error; err != nil {
			return 0, err
	}
	jumlahPinjaman := pinjaman.Amount
	sukuBunga := pinjaman.Loan_Interest_Rates
	durasiPinjaman := pinjaman.Loan_Duration
	pembayaranPerBulan := (jumlahPinjaman * sukuBunga) / (12 * 100) + (jumlahPinjaman / durasiPinjaman)
	totalPembayaran := pembayaranPerBulan * durasiPinjaman
	return totalPembayaran, nil
}
