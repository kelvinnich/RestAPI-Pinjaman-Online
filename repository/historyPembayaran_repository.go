package repository

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type HistoryPembayaranRepository interface{
	CreateHistoryPembayaranRepository(history *model.Master_Payment_History)error
	UpdateHistoryPembayaranRepository(id int, history *model.Master_Payment_History)error
	GetAllHistoryPembayaranRepository()  ([]*model.Master_Payment_History, error)
	GetHistoryPembayaranByIdRepository(id int)(*model.Master_Payment_History, error)
	DeleteHistoryPembayaranRepository(id int )error
	GetRiwayatPembayaranNasabah(nasabahID int, status string) ([]*model.Master_Payment_History, error)
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

func(db *historyPembayaranConnection)UpdateHistoryPembayaranRepository(id int, history *model.Master_Payment_History)error{
	if err := db.db.Where("id = $1", id).Updates(history).Error; err != nil {
		return err
	}
	return nil
}

func(db *historyPembayaranConnection)GetAllHistoryPembayaranRepository()  ([]*model.Master_Payment_History, error){
	var historyPembayaran []*model.Master_Payment_History
	err := db.db.Table("Master_Payment_Historys").
			Select("Master_Payment_Historys.*, master_loans.customer_id, transaction_payment_loans.payment_status").
			Joins("JOIN transaction_payment_loans ON transaction_payment_loans.id = Master_Payment_Historys.payment_id").
			Joins("JOIN master_loans ON master_loans.id = Master_Payment_Historys.loan_id").
			Scan(&historyPembayaran).Error
	if err != nil {
			return nil, err
	}
	return historyPembayaran, nil
}

func(db *historyPembayaranConnection)GetHistoryPembayaranByIdRepository(id int)(*model.Master_Payment_History, error){
	var history model.Master_Payment_History
	if err := db.db.First(&history, id).Error; err != nil {
		return nil,err
	}
	return &history,nil
}

func(db *historyPembayaranConnection)DeleteHistoryPembayaranRepository(id int )error{
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Payment_History{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *historyPembayaranConnection) GetRiwayatPembayaranNasabah(nasabahID int, status string) ([]*model.Master_Payment_History, error) {
	var riwayatPembayaran []*model.Master_Payment_History
	err := db.db.Table("Master_Payment_Historys").
			Select("Master_Payment_Historys.*").
			Joins("JOIN transaction_payment_loans ON transaction_payment_loans.id = Master_Payment_Historys.payment_id").
			Joins("JOIN master_loans ON master_loans.id = Master_Payment_Historys.loan_id").
			Where("master_loans.customer_id = ? AND transaction_payment_loans.payment_status = ?", nasabahID, status).
			Scan(&riwayatPembayaran).Error
	if err != nil {
			return nil, err
	}
	return riwayatPembayaran, nil
}
