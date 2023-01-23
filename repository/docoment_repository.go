package repository 

import (
	"pinjaman-online/model"

	"gorm.io/gorm"
)

type DocumentNasabahRepository interface{
	Create(dokumen *model.Dokumen_nasabah) error
	FindByID(id int) (*model.Dokumen_nasabah, error) 
	Update(id int, dokumen *model.Dokumen_nasabah) error
	Delete(id int) error
}

type documentConnection struct {
	DB *gorm.DB
}

func NewDocumentRepository(db *gorm.DB)DocumentNasabahRepository{
	return &documentConnection{
		DB: db,
	}
}

func (r *documentConnection) Create(dokumen *model.Dokumen_nasabah) error {
	tx := r.DB.Begin()

	
	if err := tx.Create(dokumen).Error; err != nil {
			tx.Rollback()
			return err
	}

	
	if err := tx.Model(&model.Nasabah{}).Where("id = $1", dokumen.IdNasabah).Update("status_verified", true).Error; err != nil {
			tx.Rollback()
			return err
	}

	
	tx.Commit()
	return nil
}


func (r *documentConnection) FindByID(id int) (*model.Dokumen_nasabah, error) {
	var dokumen model.Dokumen_nasabah
	if err := r.DB.First(&dokumen, id).Error; err != nil {
		return nil, err
	}
	return &dokumen, nil
}

func (r *documentConnection) Update(id int, dokumen *model.Dokumen_nasabah) error {
	if err := r.DB.Model(&model.Dokumen_nasabah{}).Where("id = $1", id).Updates(dokumen).Error; err != nil {
		return err
	}
	return nil
}

func (r *documentConnection) Delete(id int) error {
	if err := r.DB.Where("id = $1", id).Delete(&model.Dokumen_nasabah{}).Error; err != nil {
		return err
	}
	return nil
}

// func(r *documentConnection) init(){
// 	r.DB.Exec("CREATE TRIGGER update_status_verified AFTER INSERT ON dokumen_nasabahs FOR EACH ROW BEGIN UPDATE nasabahs SET status_verified = true WHERE id = NEW.id_nasabah; END;")

// }