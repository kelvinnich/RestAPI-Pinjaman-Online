package model

type Master_Document_Customer struct {
	Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Customer_Id uint64 `gorm:"not null " json:"-"`
	Customer Master_Customer `gorm:"foreignKey:Customer_Id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer_id"`
	DocumentType string `gorm:"type:varchar(255)" json:"document_type"`
	FilePath string `gorm:"type:varchar(255)" json:"file_path"`
}
