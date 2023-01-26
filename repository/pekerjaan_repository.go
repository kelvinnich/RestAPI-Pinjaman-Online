package repository

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)


type RepositoryCustomerWork interface{
	AddCustomerJobs(jobs *model.Master_Jobs_customers) (*model.Master_Jobs_customers,error)
	CustomerJobUpdates(id int,jobs *model.Master_Jobs_customers) (*model.Master_Jobs_customers,error)
	SearchForCustomerJobsById(id int) (*model.Master_Jobs_customers, error)
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


func(db *connectionCustomerWork) AddCustomerJobs(jobs *model.Master_Jobs_customers) (*model.Master_Jobs_customers,error){
	if err := db.db.Create(jobs).Error; err != nil {
		return nil,err
	}

	return jobs,nil
}

func(db *connectionCustomerWork)	CustomerJobUpdates(id int,jobs *model.Master_Jobs_customers) (*model.Master_Jobs_customers,error){
	if err := db.db.Model(&model.Master_Jobs_customers{ID: id}).Updates(jobs).Error; err != nil {
		return nil,err
	}
	return jobs,nil
}

func(db *connectionCustomerWork)SearchForCustomerJobsById(id int) (*model.Master_Jobs_customers, error){
	var jobs model.Master_Jobs_customers
	if err := db.db.First(&jobs, id).Error; err != nil {
		return nil,err
	}

	return &jobs,nil
}

func(db *connectionCustomerWork)DeleteCustomerJobs(id int)error{
	if err := db.db.Where("id = $1", id).Delete(&model.Master_Jobs_customers{}).Error; err != nil{
		return err
	}
	return nil
}