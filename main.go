package main

import (
	"pinjaman-online/config"
	"pinjaman-online/controller"
	"pinjaman-online/repository"
	"pinjaman-online/service"

	"github.com/gin-gonic/gin"
)
var(
	db = config.ConnectDB()

	//repository
	nasabahRepository repository.NasabahRepository = repository.NewNasabahRepository(db)

	//service
	jwtService service.JwtService = service.NewJwtService()
	authService service.AuthenticationService = service.NewAuthenticationService(nasabahRepository)

	//controller
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main(){
	defer config.CloseDB(db)

	r := gin.Default()

	auth := r.Group("pinjol/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	r.Run(":3000")
}