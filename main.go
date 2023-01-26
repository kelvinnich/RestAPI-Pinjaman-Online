package main

import (
	"pinjaman-online/config"
	"pinjaman-online/controller"
	"pinjaman-online/repository"
	"pinjaman-online/service"
	"pinjaman-online/middleware"

	"github.com/gin-gonic/gin"
)
var(
	db = config.ConnectDB()

	//repository
	nasabahRepository repository.NasabahRepository = repository.NewNasabahRepository(db)
	documentRepository repository.DocumentNasabahRepository = repository.NewDocumentRepository(db)
	//service
	jwtService service.JwtService = service.NewJwtService()
	authService service.AuthenticationService = service.NewAuthenticationService(nasabahRepository)
	nasabahService service.NasabahServic = service.NewNasabahService(nasabahRepository)
	documentService service.DocumentService = service.NewDocumentService(documentRepository)

	
	//controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	nasabahController controller.NasabahController = controller.NewNasabahController(nasabahService, jwtService)
	documentController controller.DocumentNasabahController = controller.NewDocumentController(documentService, jwtService)
)

func main(){
	defer config.CloseDB(db)

	r := gin.Default()

	auth := r.Group("pinjol/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	nasabah := r.Group("pinjol/nasabah", middleware.Authorize(jwtService))
	{
		nasabah.PUT("/update", nasabahController.UpdateNasabahController)
		nasabah.GET("/profile", nasabahController.ProfileNasabahController)
	}

	documentNasabah := r.Group("pinjol/document", middleware.Authorize(jwtService))
	{
		documentNasabah.POST("/uploadDocs", documentController.UploadDocumentController)
		documentNasabah.PUT("/:id", documentController.UpdateDocumentController)
		documentNasabah.GET("/:id", documentController.FindDocumentByIdController)
		documentNasabah.DELETE("/:id", documentController.DeleteDocumentController)
	}

	r.Run(":3000")
}