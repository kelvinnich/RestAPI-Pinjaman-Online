package main

import (
	"log"
	"pinjaman-online/config"
	"pinjaman-online/controller"
	"pinjaman-online/middleware"
	"pinjaman-online/repository"
	"pinjaman-online/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
var(
	db *gorm.DB = config.ConnectDB()
	

	//repository
	nasabahRepository repository.NasabahRepository = repository.NewNasabahRepository(db)
	documentRepository repository.DocumentNasabahRepository = repository.NewDocumentRepository(db)
	pekerjaanRepository repository.RepositoryCustomerWork =  repository.NewRepositoryCustomerWork(db)
	pinjamanRepository repository.PinjamanRepository = repository.NewPinjamanRepository(db)
	pembayaranRepository repository.PembayaranRepository = repository.NewPembayaranRepository(db)
	historyPembayaran repository.HistoryPembayaranRepository = repository.NewHistoryPembayaranRepository(db)

	//service
	jwtService service.JwtService = service.NewJwtService()
	authService service.AuthenticationService = service.NewAuthenticationService(nasabahRepository)
	nasabahService service.NasabahServic = service.NewNasabahService(nasabahRepository)
	documentService service.DocumentService = service.NewDocumentService(documentRepository)
	pekerjaanNasabahService service.PekerjaanNasabahService = service.NewPekerjaanNasabahService(pekerjaanRepository)
	pinjamanService service.PinjamanService = service.NewPinjamanService(pinjamanRepository,nasabahRepository)
	pembayaranService service.PembayaranService = service.NewPembayaranService(pembayaranRepository)
	historyService service.HistoryPembayaranService = service.NewHistoryPembayaranService(historyPembayaran)
	
	//controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	nasabahController controller.NasabahController = controller.NewNasabahController(nasabahService, jwtService)
	documentController controller.DocumentNasabahController = controller.NewDocumentController(documentService, jwtService)
	pekerjaanNasabahController controller.PekerjaanNasabahController = controller.NewPekerjaanNasabahController(pekerjaanNasabahService,jwtService)
	pinjamanController controller.PinjamanController = controller.NewPinajamanController(pinjamanService, jwtService)
	pembayaranController controller.PembayaranController = controller.NewPembayaranController(pembayaranService, jwtService)
	historyPembayaranController controller.HistoryPembayaranController = controller.NewHistoryPembayaranController(historyService, jwtService)
)

func main(){
	defer config.CloseDB(db)
	log.Println(db)

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

	pekerjaanNasabah := r.Group("pinjol/pekerjaan", middleware.Authorize(jwtService))
	{
		pekerjaanNasabah.POST("/addJobs", pekerjaanNasabahController.AddCustomerJobsController)
		pekerjaanNasabah.PUT("/:id", pekerjaanNasabahController.CustomerUpdateJobsController)
		pekerjaanNasabah.GET("/:id", pekerjaanNasabahController.SearchForCustomerJobsByIdController)
		pekerjaanNasabah.DELETE("/:id", pekerjaanNasabahController.DeleteCustomerJobsController)
	}

	pinjamanNasabah := r.Group("pinjol/pinjaman", middleware.Authorize(jwtService))
	{
		pinjamanNasabah.POST("/uang", pinjamanController.CreatePinjamanController)
		pinjamanNasabah.PUT("/:id", pinjamanController.UpdatePinjamanController)
		pinjamanNasabah.GET("/:id", pinjamanController.SearchPinjamanByIdController)
		pinjamanNasabah.GET("/verifikasi/:id", pinjamanController.UpdateStatusApprovalPinjamanController)
		pinjamanNasabah.DELETE("/:id", pinjamanController.DeletePinjamanController)
	}

	pembayaranNasabah := r.Group("pinjol/pembayaran", middleware.Authorize(jwtService))
	{
		pembayaranNasabah.POST("/perbulan", pembayaranController.PembayaranPinjamanController)
		pembayaranNasabah.GET("/status/:status", pembayaranController.ListPembayaranByStatusController)
		pembayaranNasabah.PUT("/:id", pembayaranController.UpdatePembayaranController)
		pembayaranNasabah.GET("/:id", pembayaranController.GetPembayaranPerBulanController)
		pembayaranNasabah.GET("/totalPembayaran/:id", pembayaranController.GetTotalPembayaranController)
		pembayaranNasabah.DELETE("/:id", pembayaranController.DeletePembayaranController)
	}

	historyPembayaranNasabah := r.Group("history/pembayaran", middleware.Authorize(jwtService))
	{
		historyPembayaranNasabah.GET("/", historyPembayaranController.GetAllHistoryPembayaranController)
		historyPembayaranNasabah.GET("/:id", historyPembayaranController.GetHistoryPembayaranByIdController)
		historyPembayaranNasabah.PUT("/:id", historyPembayaranController.UpdateHistoryPembayaranController)
	historyPembayaranNasabah.DELETE(":id", historyPembayaranController.DeleteHistoryPembayaranController)
	}

	r.Run(":3000")
}