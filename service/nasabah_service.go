package service

import (
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"

	"github.com/mashingan/smapping"
)


type NasabahServic interface{
	UpdateNasabah(nasabah dto.UpdateNasabahDTO)model.Master_Customer
	ProfileNasabah(nasabahId string)model.Master_Customer
}

type nasabahService struct{
	nasabahRepository repository.NasabahRepository
}

func NewNasabahService(nasabahRepo repository.NasabahRepository)NasabahServic{
	return &nasabahService{
		nasabahRepository: nasabahRepo,
	}
}


func(s *nasabahService)UpdateNasabah(nasabah dto.UpdateNasabahDTO)model.Master_Customer{
	NewNasabah := model.Master_Customer{}
	err := smapping.FillStruct(&NewNasabah, smapping.MapFields(nasabah))
	if err != nil {
		log.Println("Error map %v", err)
	}
	update,_ := s.nasabahRepository.UpdateNasabah(nasabah.Id, NewNasabah)
	return update
}

func(s *nasabahService)ProfileNasabah(nasabahId string)model.Master_Customer{
	return s.nasabahRepository.ProfileNasabah(nasabahId)
}