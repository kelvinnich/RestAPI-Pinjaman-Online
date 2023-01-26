package repository

import (
	"errors"
	
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type PinjamanRepository interface{
	CreatePinjamanRepository(pinjaman *model.Master_Loan) error
	UpdatePinjamanRepository(id uint64, pinjaman *model.Master_Loan)error
	SearchPinjamanByIdRepository(id uint64) (*model.Master_Loan, error)
	DeletePinjamanRepository(id uint64)error
	UpdateLoanStatus( customerID uint64) (*model.Master_Loan, error)
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

func(db *pinjamanConnection)SearchPinjamanByIdRepository(id uint64) (*model.Master_Loan, error){
	var pinjaman model.Master_Loan
	if err := db.db.First(&pinjaman, id).Error; err != nil{
		return nil,err
	}

	return &pinjaman,nil
}

func(db *pinjamanConnection)DeletePinjamanRepository(id uint64)error{
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Loan{}).Error; err != nil{
		return err
	}

	return nil
}

func(db *pinjamanConnection) UpdateLoanStatus( customerID uint64) (*model.Master_Loan, error) {
	var customer model.Master_Customer
	if err := db.db.Where("id = $1", customerID).First(&customer).Error; err != nil {
		return nil, err
	}

	if customer.StatusVerified {
		var loan model.Master_Loan
		if err := db.db.Where("customer_id = $1", customerID).First(&loan).Error; err != nil {
			return nil, err
		}

		loan.StatusApproved = true
		if err := db.db.Save(&loan).Error; err != nil {
			return nil, err
		}
		return &loan, nil
	}
	return nil, errors.New("customer status not verified")
}
