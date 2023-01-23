package repository

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type HistoryPembayaranRepository interface{
	CreateHistoryPembayaranRepository(history *model.HistoryPembayaran)error
	UpdateHistoryPembayaranRepository(id int, history *model.HistoryPembayaran)error
	GetAllHistoryPembayaranRepository()  ([]*model.HistoryPembayaran, error)
	GetHistoryPembayaranByIdRepository(id int)(*model.HistoryPembayaran, error)
	DeleteHistoryPembayaranRepository(id int )error
	GetRiwayatPembayaranNasabah(nasabahID int, status string) ([]*model.HistoryPembayaran, error)
}

type historyPembayaranConnection struct{
	db *gorm.DB
}

func NewHistoryPembayaranRepository(db *gorm.DB)HistoryPembayaranRepository{
	return &historyPembayaranConnection{
		db: db,
	}
}

func(db *historyPembayaranConnection)CreateHistoryPembayaranRepository(history *model.HistoryPembayaran)error{
	if err := db.db.Create(history).Error; err != nil {
		return err
	}
	return nil
}

func(db *historyPembayaranConnection)UpdateHistoryPembayaranRepository(id int, history *model.HistoryPembayaran)error{
	if err := db.db.Where("id = $1", id).Updates(history).Error; err != nil {
		return err
	}
	return nil
}

func(db *historyPembayaranConnection)GetAllHistoryPembayaranRepository()  ([]*model.HistoryPembayaran, error){
	var history []*model.HistoryPembayaran
	if err := db.db.Find(&history).Error; err != nil {
		return nil,err
	}
	return history,nil
}

func(db *historyPembayaranConnection)GetHistoryPembayaranByIdRepository(id int)(*model.HistoryPembayaran, error){
	var history model.HistoryPembayaran
	if err := db.db.First(&history, id).Error; err != nil {
		return nil,err
	}
	return &history,nil
}

func(db *historyPembayaranConnection)DeleteHistoryPembayaranRepository(id int )error{
	if err := db.db.Where("id = $1", id).Delete(&model.HistoryPembayaran{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *historyPembayaranConnection) GetRiwayatPembayaranNasabah(nasabahID int, status string) ([]*model.HistoryPembayaran, error) {
	var riwayatPembayaran []*model.HistoryPembayaran
	err := db.db.Table("history_pembayarans").
			Select("history_pembayarans.*").
			Joins("JOIN pembayarans ON pembayarans.id = history_pembayarans.pembayaran_id").
			Joins("JOIN pinjamen ON pinjamen.id = history_pembayarans.pinjaman_id").
			Where("pinjamans.nasabah_id = ? AND pembayarans.status_pembayaran = ?", nasabahID, status).
			Scan(&riwayatPembayaran).Error
	if err != nil {
			return nil, err
	}
	return riwayatPembayaran, nil
}
