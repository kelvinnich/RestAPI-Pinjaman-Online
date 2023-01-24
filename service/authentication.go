package service

import (
	"fmt"
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)


type AuthenticationService interface{
	CreateNasabah(nasabah dto.RegisterNasabahDTO) model.Nasabah
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	IsDuplicateNIk(noKtp string) bool
	FindByNIK(nik string)model.Nasabah
}

type authenticationService struct{
	nasabahRepository repository.NasabahRepository
}

func NewAuthenticationService(nasabah repository.NasabahRepository)AuthenticationService{
	return &authenticationService{
		nasabahRepository: nasabah,
	}
}

func(s *authenticationService) CreateNasabah(nasabah dto.RegisterNasabahDTO) model.Nasabah{
	NewNasabah := model.Nasabah{}
	err := smapping.FillStruct(&NewNasabah, smapping.MapFields(nasabah))
	if err != nil {
		fmt.Errorf("Error map %v", err)
	}

	response,_ := s.nasabahRepository.CreateNasabah(NewNasabah)
	return response
}

func(s *authenticationService)VerifyCredential(email string, password string) interface{}{
	response := s.nasabahRepository.VerifyCredential(email, password)
	if v, ok := response.(model.Nasabah); ok{
		comparePassword := hashAndComparePassword(v.Password, []byte(password))
		if v.Email == email && comparePassword{
			return response
		}
		return false
	}
	return nil
}

func(s *authenticationService) IsDuplicateEmail(email string) bool{
	response := s.nasabahRepository.IsDuplicateEmail(email)
	return !(response.Error == nil)
}

func(s *authenticationService) IsDuplicateNIk(noKtp string) bool{
	response := s.nasabahRepository.IsDuplicateNIk(noKtp)
	return !(response.Error == nil)
}

func(s *authenticationService)FindByNIK(nik string)model.Nasabah{
	return s.nasabahRepository.FindByNIK(nik)
}

func hashAndComparePassword(hashPassword string, planPassword []byte)bool{
	hash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hash, planPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}