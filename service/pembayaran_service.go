package service

import (
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"
	"strconv"
	"time"

	"github.com/mashingan/smapping"
)

type PembayaranService interface {
	PembayaranPinjamanService(payment dto.CreatePembayaranDTO) (*model.Transactions_Payment_Loan, error)
	ListPembayaranByStatusService(status string) ([]*model.Transactions_Payment_Loan, error)
	UpdatePembayaranService(updatePayment dto.UpdatePembayaranDTO) (*model.Transactions_Payment_Loan, error)
}

type pembayaranService struct {
	pembayaranRepository repository.PembayaranRepository
}

func NewPembayaranService(pembayaranRepo repository.PembayaranRepository)PembayaranService{
	return &pembayaranService{
		pembayaranRepository: pembayaranRepo,
	}
}

func (s *pembayaranService) PembayaranPinjamanService(dtoPayment dto.CreatePembayaranDTO) (*model.Transactions_Payment_Loan, error) {
	var txPinjaman model.Transactions_Payment_Loan
	err := smapping.FillStruct(&txPinjaman, smapping.MapFields(&dtoPayment))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	currentMonth := time.Now().Month()
	if currentMonth == txPinjaman.Payment_Date.Month() {
		txPinjaman.Payment_Status = true
	}

	pembayaran, err := s.pembayaranRepository.CreatePembayaranRepository(&txPinjaman)
	if err != nil {
		return nil, err
	}
		
    return pembayaran, nil
}

func(s *pembayaranService) UpdatePembayaranService(updatePayment dto.UpdatePembayaranDTO) (*model.Transactions_Payment_Loan, error){
	var txPinjaman model.Transactions_Payment_Loan
	err := smapping.FillStruct(&txPinjaman, smapping.MapFields(&updatePayment))
	if err != nil {
		log.Printf("Error Map %v", err)
	}

	updateTx,err := s.pembayaranRepository.UpdatePembayaranRepository(txPinjaman.ID, &txPinjaman)
	if err != nil {
		log.Printf("Error to update pembayaran %v", err)
	}

	return updateTx, nil
}

func (s *pembayaranService) ListPembayaranByStatusService(status string) ([]*model.Transactions_Payment_Loan, error) {
	pembayarans, err := s.pembayaranRepository.ListPembayaranRepository()
	if err != nil {
			return nil, err
	}
	var filteredPembayarans []*model.Transactions_Payment_Loan
	for _, p := range pembayarans {
			if strconv.FormatBool(p.Payment_Status) == status {
					filteredPembayarans = append(filteredPembayarans, p)
			}
	}
	return filteredPembayarans, nil
}


	
