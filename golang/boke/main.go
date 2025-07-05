package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/test/boke/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 定义用户表
type User struct {
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}

// Post 定义文章表
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	UserID    uint      `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
}

// Post 定义评论表
type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"column:content"`
	UserID    uint      `gorm:"column:user_id"`
	PostID    uint      `gorm:"column:post_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type Books struct {
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"column:title"`
	Author string  `gorm:"column:author"`
	Price  float64 `gorm:"column:price"`
}

// 定义 Claims 结构体
type Claims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

// 2. 预定义几种常见错误
var (
	ErrDB         = &APIError{Code: http.StatusInternalServerError, Message: "数据库连接失败"}
	ErrAuth       = &APIError{Code: http.StatusUnauthorized, Message: "认证失败，请先登录"}
	ErrNotFound   = &APIError{Code: http.StatusNotFound, Message: "资源不存在"}
	ErrBadRequest = &APIError{Code: http.StatusBadRequest, Message: "请求参数错误"}
)

var jwtKey = []byte("my_secret_key")

// var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(&User{}, &Post{}) //创建表
	err1 := db.AutoMigrate(&Comment{})      //创建表
	if err != nil {
		panic(err)
	}
	if err1 != nil {
		panic(err1)
	}
	return db
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	// db, err := gorm.Open(sqlite.Open(constant.DBPATH), &gorm.Config{})
	dsn := "user:pass@tcp(127.0.0.1:3306)/boke?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println(err.Error())
		panic(err)
	}
	return db
}

func GenerateToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-jwt-demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 验证 JWT
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// setupAuthMiddleware 设置认证中间件
func setupAuthMiddleware() gin.HandlerFunc {
	fmt.Println(">>> JWTAuthMiddleware entered")
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		/*
			// 一般是 Bearer token
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
				c.Abort()
				return
			}

			claims, err := ValidateToken(parts[1])
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
				c.Abort()
				return
			}

			// 验证通过，把用户信息放上下文
			c.Set("username", claims.Username)

			fmt.Println(">>> JWTAuthMiddleware entered1")
			c.Next()
		*/

		// authHeader = authHeader[len("Bearer "):]

		token, err := jwt.ParseWithClaims(authHeader, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("user_id", claims.UserId)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}

func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World2!",
		})
	})

	// 鉴权路由组
	// auth := router.Group("/api")
	// auth.Use(JWTAuthMiddleware())

	//注册路由
	authMiddleware := setupAuthMiddleware()
	setupRoutes(router, db, authMiddleware)

	// auth.GET("/createPost", func(c *gin.Context) {
	// 	title := c.GetString("title")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Welcome " + title,
	// 	})
	// })

	router.Run(":8081")

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
