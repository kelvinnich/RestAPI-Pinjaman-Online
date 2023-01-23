package repository

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type RepositoryCustomerWork interface{
	AddCustomerJobs(jobs *model.Pekerjaan_nasabah) error
	CustomerJobUpdates(id int,jobs *model.Pekerjaan_nasabah) error
	SearchForCustomerJobsById(id int) (*model.Pekerjaan_nasabah, error)
	DeleteCustomerJobs(id int)error
}

type connectionCustomerWork struct{
	db *gorm.DB
}

func NewRepositoryCustomerWork(db *gorm.DB) RepositoryCustomerWork{
	return &connectionCustomerWork{
		db: db,
	}
}


func(db *connectionCustomerWork) AddCustomerJobs(jobs *model.Pekerjaan_nasabah) error{
	if err := db.db.Create(jobs).Error; err != nil {
		return err
	}

	return nil
}

func(db *connectionCustomerWork)	CustomerJobUpdates(id int,jobs *model.Pekerjaan_nasabah) error{
	if err := db.db.Model(&model.Pekerjaan_nasabah{}).Where("id = $1", id).Updates(jobs).Error; err != nil {
		return err
	}
	return nil
}

func(db *connectionCustomerWork)SearchForCustomerJobsById(id int) (*model.Pekerjaan_nasabah, error){
	var jobs model.Pekerjaan_nasabah
	if err := db.db.First(&jobs, id).Error; err != nil {
		return nil,err
	}

	return &jobs,nil
}

func(db *connectionCustomerWork)DeleteCustomerJobs(id int)error{
	if err := db.db.Where("id = $1", id).Delete(&model.Pekerjaan_nasabah{}).Error; err != nil{
		return err
	}
	return nil
}