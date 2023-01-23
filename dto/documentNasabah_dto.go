package dto

type CreateDocumentNasabahDTO struct{
	Nasabah_id uint64 `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	JenisDokumen string `json:"jenis_dokumen" form:"jenis_dokumen" binding:"required"`
	FilePath string `json:"file_path" form:"file_path" binding:"required"`
}

type UpdateDocumentNasabahDTO struct{
	Id uint64 `json:"id" form:"id"`
	Nasabah_id uint64 `json:"nasabah_id" form:"nasabah_id" binding:"required"`
	JenisDokumen string `json:"jenis_dokumen" form:"jenis_dokumen" binding:"required"`
	FilePath string `json:"file_path" form:"file_path" binding:"required"`
}