package repository

import (
	"fmt"
	"pinjaman-online/model"

	"gorm.io/gorm"
)

type DocumentNasabahRepository interface{
	Create(dokumen *model.Master_Document_Customer)( *model.Master_Document_Customer, error)
	FindByID(id uint64) (*model.Master_Document_Customer, error) 
	Update(id uint64, dokumen *model.Master_Document_Customer) (*model.Master_Document_Customer,error)
	Delete(id uint64) error
}

type documentConnection struct {
	DB *gorm.DB
}

func NewDocumentRepository(db *gorm.DB)DocumentNasabahRepository{
	return &documentConnection{
		DB: db,
	}
}

func (r *documentConnection) Create(dokumen *model.Master_Document_Customer) (*model.Master_Document_Customer, error) {
	
	tx := r.DB.Begin()

	if err := tx.Create(dokumen).Error; err != nil {
			tx.Rollback()
			return nil, err
	}

	if err := tx.Model(&model.Master_Customer{Id: dokumen.Customer_Id}).UpdateColumn("status_verified", true).Error; err != nil {
		tx.Rollback()
		return nil,err
	}
	// if err := tx.Exec("UPDATE master_customers SET status_verified = true WHERE id = $1", dokumen.Customer_Id).Error; err != nil {
	// 		tx.Rollback()
	// 		return nil, err
	// }

	tx.Commit()

	return dokumen, nil
}



func (r *documentConnection) FindByID(id uint64) (*model.Master_Document_Customer, error) {
	var dokumen *model.Master_Document_Customer
	if err := r.DB.First(&dokumen, id).Error; err != nil {
		return nil,err
	}
	return dokumen, nil
}

func (r *documentConnection) Update(id uint64, dokumen *model.Master_Document_Customer) (*model.Master_Document_Customer,error) {
	if err := r.DB.Model(&model.Master_Document_Customer{Id: id}).Updates(dokumen).Error; err != nil {
		return nil,err
	}
	fmt.Printf("documentRepo %s", dokumen)
	return dokumen,nil
}

func (r *documentConnection) Delete(id uint64) error {
	if err := r.DB.Where("id = $1", id).Delete(&model.Master_Document_Customer{}).Error; err != nil {
		return err
	}
	return nil
}

// func(r *documentConnection) init(){
// 	r.DB.Exec("CREATE TRIGGER update_status_verified AFTER INSERT ON dokumen_nasabahs FOR EACH ROW BEGIN UPDATE nasabahs SET status_verified = true WHERE id = NEW.id_nasabah; END;")

// }