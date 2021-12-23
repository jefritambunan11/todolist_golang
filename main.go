package main

import(
	"bwastartup/user"
	"bwastartup/campaign"

	"bwastartup/handler"
	"bwastartup/auth"
	"bwastartup/helper"
	
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
	
	//"fmt"
	"log"
	"strings"
	"net/http"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	userRepository := user.NewRepository(db) 
	campaignRepository := campaign.NewRepository(db) 

	userService := user.NewService(userRepository) 
	campaignService := campaign.NewService(campaignRepository) 
	authService := auth.NewService() 

	

		
	userHandler := handler.NewUserHandler(userService, authService)	
	campaignHandler := handler.NewCampaignHandler(campaignService)	



	
	router := gin.Default()
	router.Static("/images", "./images")
	api_v1 := router.Group("/api/v1")
	
	api_v1.POST("/users", userHandler.RegisterUser)
	api_v1.POST("/sessions", userHandler.Login)
	api_v1.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api_v1.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)


	api_v1.GET("/campaigns", campaignHandler.GetCampaigns)	
	api_v1.GET("/campaigns/:id", campaignHandler.GetCampaign)	
	api_v1.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)	
		


	router.Run(":5000")
	
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		
		var tokenString string

		// convert token ke array
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		_token_, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := _token_.Claims.(jwt.MapClaims)

		if !ok || !_token_.Valid {
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
