package repository

import (
	"errors"
	"log"
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type HistoryPembayaranRepository interface{
	CreateHistoryPembayaranRepository(history *model.Master_Payment_History)error
	UpdateHistoryPembayaranRepository(id uint64, history *model.Master_Payment_History)(model.Master_Payment_History,error)
	GetAllHistoryPembayaranRepository()  ([]*model.Master_Payment_History, error)
	GetHistoryPembayaranByIdRepository(id uint64)(*model.Master_Payment_History, error)
	DeleteHistoryPembayaranRepository(id uint64 )error
	GetRiwayatPembayaranNasabahById(nasabahID uint64) ([]*model.Master_Payment_History, error)
}

type historyPembayaranConnection struct{
	db *gorm.DB
}

func NewHistoryPembayaranRepository(db *gorm.DB)HistoryPembayaranRepository{
	return &historyPembayaranConnection{
		db: db,
	}
}

func(db *historyPembayaranConnection)CreateHistoryPembayaranRepository(history *model.Master_Payment_History)error{
	if err := db.db.Create(history).Error; err != nil {
		return err
	}
	return nil
}

func(db *historyPembayaranConnection)UpdateHistoryPembayaranRepository(id uint64, history *model.Master_Payment_History)(model.Master_Payment_History,error){
	if err := db.db.Model(&model.Master_Payment_History{Id: id}).Updates(history).Error; err != nil {
		log.Printf("failed tp update history pembayaran repository %v",err)
	}
	return *history,nil
}

func(db *historyPembayaranConnection)GetAllHistoryPembayaranRepository()  ([]*model.Master_Payment_History, error){
	var historyPembayaran []*model.Master_Payment_History
	err := db.db.Table("master_payment_histories").
			Select("master_payment_histories.*, master_loans.customer_id, transactions_payment_loans.payment_status").
			Joins("JOIN transactions_payment_loans ON transactions_payment_loans.id = master_payment_histories.payment_id").
			Joins("JOIN master_loans ON master_loans.id = master_payment_histories.loan_id").
			Scan(&historyPembayaran).Error
			if err != nil {
				log.Println(err)
				return nil, err
			}
			if len(historyPembayaran) == 0 {
				log.Println("Data tidak ditemukan")
				return nil, errors.New("Data tidak ditemukan")
			}
	return historyPembayaran, nil
}

func(db *historyPembayaranConnection)GetHistoryPembayaranByIdRepository(id uint64)(*model.Master_Payment_History, error){
	var history model.Master_Payment_History
	if err := db.db.First(&history, id).Error; err != nil {
		return nil,err
	}
	return &history,nil
}

func(db *historyPembayaranConnection)DeleteHistoryPembayaranRepository(id uint64 )error{
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Payment_History{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *historyPembayaranConnection) GetRiwayatPembayaranNasabahById(id uint64) ([]*model.Master_Payment_History, error) {
		var historyPembayaran []*model.Master_Payment_History
		err := db.db.Table("master_payment_histories").
		Select("master_payment_histories.*, master_loans.customer_id, transactions_payment_loans.payment_status").
		Joins("JOIN transactions_payment_loans ON transactions_payment_loans.id = master_payment_histories.payment_id").
		Joins("JOIN master_loans ON master_loans.id = master_payment_histories.loan_id").
		Where("master_payment_histories.id = $1", id).
		Scan(&historyPembayaran).Error
		if err != nil {
		log.Println(err)
		return nil, err
		}
		if historyPembayaran[0].Id == 0 {
		log.Println("Data tidak ditemukan")
		return nil, errors.New("Data tidak ditemukan")
		}
		return historyPembayaran, nil
	}

