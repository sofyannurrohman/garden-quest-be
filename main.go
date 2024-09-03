package main

import (
	"garden-quest/auth"
	"garden-quest/handler"
	"garden-quest/helper"
	"garden-quest/plant"
	"garden-quest/transaction"
	"garden-quest/user"
	"garden-quest/water"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/garden?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Repository
	userRepository := user.NewRepository(db)
	plantRepository := plant.NewRepository(db)
	waterRepository := water.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	//Service
	plantService := plant.NewService(plantRepository)
	waterService := water.NewService(waterRepository)
	userService := user.NewService(userRepository, plantService, waterRepository, waterService)
	authService := auth.NewService()
	transactionService := transaction.NewService(transactionRepository)
	// Handler
	userHandler := handler.NewUserHandler(userService, authService)
	plantHandler := handler.NewPlantHandler(plantService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	router := gin.Default()
	// router.Use(cors.Default())

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	api.POST("/email-checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/plants", authMiddleware(authService, userService), plantHandler.GetUserPlant)
	api.POST("/watering", authMiddleware(authService, userService), userHandler.AddWater)
	api.POST("/energy/:userID", userHandler.AddEnergy)
	api.POST("/buy/plants", authMiddleware(authService, userService), userHandler.BuyPlantType)
	api.POST("/buy/waters", authMiddleware(authService, userService), userHandler.BuyWaterEnergy)

	// api.GET("/transactions", authMiddleware(authService, userService), transactionHandler)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.GET("/users/:userID/inventory", userHandler.GetInventory)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
