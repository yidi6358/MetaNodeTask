package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func registerHandler(db *gorm.DB) gin.HandlerFunc {
	// 注册接口
	return func(c *gin.Context) {
		// 1. 绑定请求数据
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误: " + err.Error(),
			})
			return
		}

		// 2. 检查用户名是否已存在
		var user User

		result := db.Model(&User{}).Where("user_name = ?", newUser.UserName).First(&user)
		if result.RowsAffected > 0 {
			c.JSON(http.StatusConflict, gin.H{
				"code":    409,
				"message": "用户名已存在",
			})
			return
		}

		// 3. 密码加密处理
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "密码加密失败",
			})
			return
		}

		// 4. 保存用户信息（实际项目中应存入数据库）
		// userDB[newUser.UserName] = User{
		// 	UserName: newUser.UserName,
		// 	Password: string(hashedPassword), // 存储加密后的密码
		// 	Email:    newUser.Email,
		// }

		result1 := db.Create(&User{UserName: newUser.UserName, Password: string(hashedPassword)})
		fmt.Printf("保存用户信息：%v\n", result1)

		// 5. 返回成功响应
		c.JSON(http.StatusCreated, gin.H{
			"code":    201,
			"message": "注册成功",
			"data": gin.H{
				"username": newUser.UserName,
				"email":    newUser.Email,
			},
		})
	}
}

func loginHandler(db *gorm.DB) gin.HandlerFunc {
	// 登录接口
	return func(c *gin.Context) {
		var loginUser struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
			return
		}

		// 从数据库获取用户(这里用模拟的userDB)
		// storedUser, exists := userDB[loginUser.Username]

		var user User
		result := db.Model(&User{}).Where("user_name = ?", loginUser.Username).First(&user)

		if result.RowsAffected == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			return
		}

		// 比较密码
		err := bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(loginUser.Password),
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
			return
		}

		token, err := GenerateToken(user.ID)

		// 登录成功
		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}
