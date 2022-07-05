package main

import (
	"bank/auth"
	"bank/handler"
	"bank/user"
	"fmt"

	"log"

	"bank/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Bank API
// @version         1.0
// @description     This is a Bank server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "Sample of Bank Server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	dsn := "host=localhost user=postgres password=postgres dbname=bank port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNX0.oYtZJsJE5SLFX907UZohtIXdQZSXs97LWgIP-s1RCNw")

	if err != nil {
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}
	userHandler := handler.NewUserHandler(userService, authService)

	fmt.Println(authService.GenerateToken(1))

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// swagger:route POST /user CreateUser
		v1.POST("/user", userHandler.RegisterUser)
		v1.POST("/login", userHandler.Login)
		v1.POST("/email_checkers", userHandler.CheckEmail)
		v1.POST("/upload_avatar", userHandler.UploadAvatar)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

}
