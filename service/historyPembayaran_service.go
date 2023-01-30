package service

import (
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"
	"time"

	"github.com/mashingan/smapping"
)

type HistoryPembayaranService interface {
	GetAllHistoryPembayaran() ([]*model.Master_Payment_History, error)
	GetHistoryPembayaranByID(id uint64) ([]*model.Master_Payment_History, error)
	UpdateHistoryPembayaranService(update dto.UpdateHistoryPembayaranDTO) (model.Master_Payment_History,error)
	DeleteHistoryPembayaranService(id uint64) error
}

type historyPembayaranService struct {
	historyRepo repository.HistoryPembayaranRepository
}

func NewHistoryPembayaranService(hsp repository.HistoryPembayaranRepository) HistoryPembayaranService{
	return &historyPembayaranService{
		historyRepo: hsp,
	}
}

  func (s *historyPembayaranService) GetAllHistoryPembayaran() ([]*model.Master_Payment_History, error) {
	histories, err := s.historyRepo.GetAllHistoryPembayaranRepository()
	if err != nil {
	return nil, err
	}
	return histories, nil
	}

	func(s *historyPembayaranService) GetHistoryPembayaranByID(id uint64) ([]*model.Master_Payment_History, error){
		historyById,err := s.historyRepo.GetRiwayatPembayaranNasabahById(id )
		if err != nil {
			log.Printf("error history service %v", err)
		}

		return historyById,nil
	}

	func (s *historyPembayaranService) UpdateHistoryPembayaranService(update dto.UpdateHistoryPembayaranDTO) (model.Master_Payment_History,error){
		var updateHistorty model.Master_Payment_History
		err := smapping.FillStruct(&updateHistorty, smapping.MapFields(&update))
		if err != nil {
			log.Printf("Error map %v", err)
		}

		updates,err := s.historyRepo.UpdateHistoryPembayaranRepository(updateHistorty.Id, &updateHistorty)
		if err != nil {
			log.Printf("failed to update history pembayaran service %v", err)
		}

		updates.Date = time.Now()

		return updates,nil
	}

	func(s *historyPembayaranService) DeleteHistoryPembayaranService(id uint64) error{
		return s.historyRepo.DeleteHistoryPembayaranRepository(id)
	}