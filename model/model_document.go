package model

type Master_Document_Customer struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Id_Customer uint `gorm:"not null " json:"-"`
	Customer_Id Master_Customer `gorm:"foreignKey:Id_Customer;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer_id"`
	DocumentType string `gorm:"type:varchar(255)" json:"document_type"`
	FilePath string `gorm:"type:varchar(255)" json:"file_path"`
}
