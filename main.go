package main

import (
	"PAK/auth"
	"PAK/handler"
	"PAK/helper"
	"PAK/komponen"
	"PAK/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pak?charset=utf8mb4&parseTime=True&loc=Local" //memasukkan hak akses
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//komponenpak
	komponenpakRepository := komponen.NewRepository(db)
	komponenpakService := komponen.NewServiceKomponenPAK(komponenpakRepository)
	komponenpakhandler := handler.NewkomponenpakHandler(komponenpakService)

	//Auto Migration DB
	//db.AutoMigrate(&user.User{})
	db.AutoMigrate(&komponen.Entity_pak{})
	//user

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("api/komponenpak")
	api2 := router.Group("api/user")

	api.POST("/create", komponenpakhandler.Create)
	api.GET("/komponen", komponenpakhandler.GetKomponen)
	api2.POST("/register", userHandler.Register)
	api2.POST("/login", userHandler.Login)
	api2.POST("/email_available", userHandler.CheckEmailAvailability)
	api2.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized 1", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Bearer tokentoken
		// memisahkan token berdasarkan spasi

		tokenString := ""
		arraytokenString := strings.Split(authHeader, " ")
		if len(arraytokenString) == 2 {
			tokenString = arraytokenString[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorized 2", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized 3", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		userID := int(claim["user_id"].(float64))
		fmt.Println(userID)

		user, err := userService.GetUserbyID(userID)

		if err != nil {
			response := helper.ApiResponse("Unauthorized 4", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
