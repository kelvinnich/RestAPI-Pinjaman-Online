package service

import (
	"log"
	"pinjaman-online/dto"
	"pinjaman-online/model"
	"pinjaman-online/repository"

	"github.com/mashingan/smapping"
)


type DocumentService interface {
	UploadDocument(document dto.CreateDocumentNasabahDTO) *model.Master_Document_Customer
	UpdateDocument(document dto.UpdateDocumentNasabahDTO) *model.Master_Document_Customer
	GetDocumentById(id uint64) (*model.Master_Document_Customer, error)
	DeleteDocument(id uint64)error
}

type documentService struct{
	documentRepository repository.DocumentNasabahRepository
}

func NewDocumentService(documentRepository repository.DocumentNasabahRepository)DocumentService{
	return &documentService{
		documentRepository: documentRepository,
	}
}

func(s *documentService) UploadDocument(document dto.CreateDocumentNasabahDTO) *model.Master_Document_Customer{
	var documents model.Master_Document_Customer
	err := smapping.FillStruct(&documents, smapping.MapFields(&document))
	if err != nil{
		log.Printf("Error map %v", err)
	}

	upload,_ := s.documentRepository.Create(&documents)
	return upload
}

func(s *documentService) UpdateDocument(document dto.UpdateDocumentNasabahDTO) *model.Master_Document_Customer{
	var updateDocument model.Master_Document_Customer
	err := smapping.FillStruct(&updateDocument, smapping.MapFields(document))
	if err != nil {
		log.Printf("Error map %v", )
	}

	update,_ := s.documentRepository.Update(updateDocument.Id, &updateDocument)
	return update
}

func(s *documentService) GetDocumentById(id uint64) (*model.Master_Document_Customer, error){
	return s.documentRepository.FindByID(id)
}

func(s *documentService) DeleteDocument(id uint64)error{
	return s.documentRepository.Delete(id)
}