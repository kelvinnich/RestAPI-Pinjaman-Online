package repository

import (
	"pinjaman-online/model"
	"time"

	"gorm.io/gorm"
)

type PembayaranRepository interface {
	CreatePembayaranRepository(pembayaran *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan,error)
	FindByIdRepository(id int) (*model.Transactions_Payment_Loan, error)
	UpdatePembayaranRepository(id int, pembayaran *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan,error)
	DeletePembayaranRepository(id int) error
	ListPembayaranRepository() ([]*model.Transactions_Payment_Loan, error)
	GetPembayaranPerBulanRepository(pinjamanID int) (int, error)
	GetTotalPembayaranRepository(pinjamanID int) (int, error)
	GetJatuhTempoPembayaranRepository(pinjamanID int) ([]time.Time, error)
}

type pembayaranConnection struct {
	db *gorm.DB
}


func NewPembayaranRepository(db *gorm.DB)PembayaranRepository{
	return &pembayaranConnection{
		db: db,
	}
}

func (db *pembayaranConnection) CreatePembayaranRepository(pembayaran *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan, error) {
	if err := db.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(pembayaran).Error; err != nil {
					return err
			}
			if err := tx.Model(&model.Master_Loan{Id: pembayaran.Loan.Id}).Where("id", pembayaran.Loan_id).UpdateColumn("amount", gorm.Expr("amount - ?", pembayaran.Monthly_Payments)).Error; err != nil {
					return err
			}
			return nil
	}); err != nil {
			return nil, err
	}
	return pembayaran, nil
}


func(db *pembayaranConnection )UpdatePembayaranRepository(id int, pembayaran *model.Transactions_Payment_Loan) (*model.Transactions_Payment_Loan,error){
	if err := db.db.Model(&model.Transactions_Payment_Loan{}).Where("id", id).Updates(pembayaran).Error; err != nil {
		return nil,err
	}

	return pembayaran,nil
}

func(db *pembayaranConnection)FindByIdRepository(id int) (*model.Transactions_Payment_Loan, error){
	var pembayaran model.Transactions_Payment_Loan
	if err := db.db.First(&pembayaran, id).Error; err != nil {
		return nil,err
	}

	return &pembayaran,nil
}

func(db *pembayaranConnection)DeletePembayaranRepository(id int) error{
	if err := db.db.Where("id = $1", id).Delete(&model.Transactions_Payment_Loan{}).Error; err != nil {
		return err
	}
	return nil
}

func(db *pembayaranConnection)ListPembayaranRepository() ([]*model.Transactions_Payment_Loan, error){
    var pembayarans []*model.Transactions_Payment_Loan
    if err := db.db.Find(&pembayarans).Error; err != nil {
        return nil, err
    }
    return pembayarans, nil
}

func (db *pembayaranConnection) GetPembayaranPerBulanRepository(pinjamanID int) (int, error) {
	var pinjaman model.Master_Loan
	if err := db.db.First(&pinjaman, pinjamanID).Error; err != nil {
			return 0, err
	}
	jumlahPinjaman := pinjaman.Amount
	sukuBunga := pinjaman.Loan_Interest_Rates
	durasiPinjaman := pinjaman.Loan_Duration
	pembayaranPerBulan := (jumlahPinjaman * sukuBunga) / (12 * 100) + (jumlahPinjaman / durasiPinjaman)
	return pembayaranPerBulan, nil
}

func (db *pembayaranConnection) GetTotalPembayaranRepository(pinjamanID int) (int, error) {
	var pinjaman model.Master_Loan
	if err := db.db.First(&pinjaman, pinjamanID).Error; err != nil {
	return 0, err
	}
	jumlahPinjaman := pinjaman.Amount
	sukuBunga := pinjaman.Loan_Interest_Rates
	durasiPinjaman := pinjaman.Loan_Duration
	pembayaranPerBulan := (jumlahPinjaman * sukuBunga) / (12 * 100) + (jumlahPinjaman / durasiPinjaman)
	totalPembayaran := pembayaranPerBulan * durasiPinjaman
	return totalPembayaran, nil
	}

func (db *pembayaranConnection) GetJatuhTempoPembayaranRepository(pinjamanID int) ([]time.Time, error) {
    var pembayarans []*model.Transactions_Payment_Loan
    var jatuhTempo []time.Time
    var pinjaman model.Master_Loan

    if err := db.db.First(&pinjaman, pinjamanID).Error; err != nil {
        return nil, err
    }

    if err := db.db.Where("loan_id = $1", pinjamanID).Find(&pembayarans).Error; err != nil {
        return nil, err
    }

    for _, pembayaran := range pembayarans {
        jatuhTempo = append(jatuhTempo, pembayaran.Payment_Date.AddDate(0, int(pinjaman.Loan_Duration), 0))
    }

    return jatuhTempo, nil
}




