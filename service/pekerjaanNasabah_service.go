package service

import (
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"

	"github.com/mashingan/smapping"
)


type PekerjaanNasabahService interface{
	AddCustomerJobsService(jobs dto.CreatePekerjaanNasabahDTO) (*model.Master_Jobs_customers, error)
	CustomersJobsUpdateService(jobs dto.UpdatePekerjaanNasabahDTO) (*model.Master_Jobs_customers,error)
	SearchForCustomerJobsByIdService(id int) (*model.Master_Jobs_customers, error)
	DeleteCustomerJobsService(id int)error
}

type pekerjaanNasabahService struct{
	jobsRepository repository.RepositoryCustomerWork
}

func NewPekerjaanNasabahService(jp repository.RepositoryCustomerWork)PekerjaanNasabahService{
	return &pekerjaanNasabahService{
		jobsRepository: jp,
	}
}

func(s *pekerjaanNasabahService)AddCustomerJobsService(jobs dto.CreatePekerjaanNasabahDTO) (*model.Master_Jobs_customers, error){
	var jobsCutomer model.Master_Jobs_customers
	err := smapping.FillStruct(&jobsCutomer, smapping.MapFields(&jobs))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	addJobs, err := s.jobsRepository.AddCustomerJobs(&jobsCutomer)
	if err != nil {
		log.Printf("error add customer %v", err)
	}

	return addJobs,nil
}

func(s *pekerjaanNasabahService)CustomersJobsUpdateService(jobs dto.UpdatePekerjaanNasabahDTO) (*model.Master_Jobs_customers,error){
	var jobsUpdate model.Master_Jobs_customers
	err := smapping.FillStruct(&jobsUpdate, smapping.MapFields(&jobs))
	if err != nil {
		log.Printf("Error map %v", err)
	}

	updateJobs, err := s.jobsRepository.CustomerJobUpdates(jobs.Id, &jobsUpdate)
	if err != nil {
		log.Printf("failed to update customer jobs %v", err)
	}

	return updateJobs,nil
}

func(s *pekerjaanNasabahService)SearchForCustomerJobsByIdService(id int) (*model.Master_Jobs_customers, error){
	return s.jobsRepository.SearchForCustomerJobsById(id)
}

func(s *pekerjaanNasabahService)DeleteCustomerJobsService(id int)error{
	return s.jobsRepository.DeleteCustomerJobs(id)
}