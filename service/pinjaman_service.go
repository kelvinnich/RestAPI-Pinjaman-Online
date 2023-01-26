package service

import (
	"errors"
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"

	"github.com/mashingan/smapping"
)


type PinjamanService interface{
	CreatePinjamanService(pinjaman dto.CreatePinjamanDTO) (*model.Master_Loan, error)
	UpdatePinjamanService(pinjaman dto.UpdatePinjamanDTO)(*model.Master_Loan, error)
	SearchPinjamanByIdService(id uint64)(*model.Master_Loan, error)
	DeletePinjamanService(id uint64)error
	UpdateLoanStatus(customerId uint64) (*model.Master_Loan, error)
}

type pinjamanService struct{
	pinjamanRepo repository.PinjamanRepository
}

func NewPinjamanService(pinjamanRepo repository.PinjamanRepository)PinjamanService{
	return &pinjamanService{
		pinjamanRepo: pinjamanRepo,
	}
}

func(s *pinjamanService)CreatePinjamanService(pinjaman dto.CreatePinjamanDTO) (*model.Master_Loan, error){
	pinjamans := &model.Master_Loan{}
	err := smapping.FillStruct(&pinjamans, smapping.MapFields(&pinjaman))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	var customer model.Master_Customer 
	

	if customer.StatusVerified{
		if pinjaman.Amount > 500000 {
			return nil, errors.New("batas peminjaman untuk user belum terverifikasi adalah 500000")
		}

		}else {
			if pinjaman.Amount > 10000000 {
				return nil, errors.New("batas peminjaman untuk user sudah terverifikasi adalah 10000000")
	}
}

	err = s.pinjamanRepo.CreatePinjamanRepository(pinjamans)
	if err != nil {
		return nil, err
	}
	return pinjamans, nil

}

func(s *pinjamanService) UpdatePinjamanService(pinjaman dto.UpdatePinjamanDTO)(*model.Master_Loan, error){
	pinjamans := &model.Master_Loan{}
	err := smapping.FillStruct(&pinjamans, smapping.MapFields(&pinjaman))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	var customer model.Master_Customer
	

	if customer.StatusVerified == false{
		if pinjaman.Amount > 500000 {
			return nil, errors.New("batas peminjaman untuk user belum terverifikasi adalah 500000")
		}

		}else if customer.StatusVerified == true{
			if pinjaman.Amount > 10000000 {
				return nil, errors.New("batas peminjaman untuk user sudah terverifikasi adalah 10000000")
	}
}


	err = s.pinjamanRepo.UpdatePinjamanRepository(pinjamans.Id, pinjamans)
	if err != nil {
		return nil, err
	}
	return pinjamans, nil
}


func(s *pinjamanService) SearchPinjamanByIdService(id uint64)(*model.Master_Loan, error){
	return s.pinjamanRepo.SearchPinjamanByIdRepository(id)
}

func(s *pinjamanService) DeletePinjamanService(id uint64)error{
	return s.pinjamanRepo.DeletePinjamanRepository(id)
}

func (s *pinjamanService) UpdateLoanStatus(customerId uint64) (*model.Master_Loan, error) {
	
	loan, err := s.pinjamanRepo.UpdateLoanStatus(customerId)
	if err != nil {
			return nil, err
	}
	return loan, nil
}

