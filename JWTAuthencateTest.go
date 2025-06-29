package main

import (
	_ "github.com/echomas/DebugWithSourceCodeGin/docs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

var DB *gorm.DB

var jwtKey = []byte("sjdlfjaqioheksabnskfjndaslsdfhksahdfas")

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
}

type UserRegistration2 struct {
	Username string `json:"username" binding:"required,min=4,max=12"`
	Password string `json:"password" binding:"required,gte=6,lte=20"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "request header authorization missed",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		tokenString := parts[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid signature",
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "invalid token",
				"details": err.Error(),
			})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

// @Summary 用户注册
// @Description 使用用户名、密码和邮箱注册一个新用户
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body RegisterInput true "注册信息"
// @Success 200 {object} map[string]string "{"message": "User registered successfully"}"
// @Failure 400 {object} map[string]string "{"error": "error message"}"
// @Router /auth/register [post]
func Register2(c *gin.Context) {
	var registrationData UserRegistration2
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := User{
		Username: registrationData.Username,
		Password: registrationData.Password,
		Email:    registrationData.Email,
	}
	result := DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user registered",
		"userID":  user.ID,
	})
}

// @Summary 用户登录
// @Description 使用用户名和密码登录，成功后返回 JWT
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body LoginInput true "登录凭证"
// @Success 200 {object} map[string]string "{"token": "jwt_token_string"}"
// @Failure 400 {object} map[string]string "{"error": "error message"}"
// @Router /auth/login [post]
func Login2(c *gin.Context) {
	var credentials LoginCredentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user User
	if err := DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "username or password incorrect",
		})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "can not generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// @Summary 获取用户个人资料
// @Description 获取当前登录用户的个人信息
// @Tags User
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]models.User
// @Failure 401 {object} map[string]string "{"error": "error message"}"
// @Router /api/profile [get]
func GetProfile2(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user id not found",
		})
		return
	}
	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"joined_at": user.CreatedAt,
	})
}

// @title Gin Auth API
// @version 1.0
// @description 这是一个使用 Gin 框架构建的带 JWT 认证的示例 API 服务。
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	dsn := "root:your_password@tcp(127.0.0.1:3306)/gin_app?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})
	router := gin.Default()

	public := router.Group("/auth")
	{
		public.POST("/register", Register2)
		public.POST("/login", Login2)
	}

	protected := router.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", GetProfile2)
	}
	router.Run(":8080")
}
