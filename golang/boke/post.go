package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func createPostHandler(db *gorm.DB) gin.HandlerFunc {
	// 创建文章
	return func(c *gin.Context) {
		var thisPost struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			UserId  uint   `json:"userid"`
		}

		if err := c.ShouldBindJSON(&thisPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
			return
		}
		db.Create(&Post{Title: thisPost.Title, Content: thisPost.Content, UserID: thisPost.UserId})

		fmt.Printf("创建文章：%v\n", thisPost)
		c.JSON(http.StatusOK, gin.H{"msg": "创建文章成功"})
	}
}

func listPostHandler(db *gorm.DB) gin.HandlerFunc {
	// 获取所有文章列表信息
	return func(c *gin.Context) {
		var posts []Post
		db.Model(&Post{}).Find(&posts)
		c.JSON(http.StatusOK, gin.H{"msg": posts})
	}
}

func updatePostHandler(db *gorm.DB) gin.HandlerFunc {
	// 更新文章信息
	return func(c *gin.Context) {
		var thisPost struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			UserId  uint   `json:"userid"`
			Id      uint   `json:"id"`
		}

		if err := c.ShouldBindJSON(&thisPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
			return
		}
		var post Post
		db.Model(&Post{}).Where("id = ?", thisPost.Id).First(&post)
		if post.UserID != thisPost.UserId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "不能更新不是自己的文章"})
			return
		}
		// db.Model(&post).Update(&Post{Title: thisPost.Title, Content: thisPost.Content}).Where("id = ?", thisPost.Id)
		db.Model(&post).Where("id = ?", thisPost.Id).Updates(map[string]interface{}{
			"title":   thisPost.Title,
			"content": thisPost.Content,
		})
		c.JSON(http.StatusOK, gin.H{"msg": "更新文章成功"})
	}
}

func deletePostHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var thisPost struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			UserId  uint   `json:"userid"`
			Id      uint   `json:"id"`
		}

		if err := c.ShouldBindJSON(&thisPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
			return
		}
		var post Post
		db.Model(&Post{}).Where("id = ?", thisPost.Id).First(&post)
		if post.UserID != thisPost.UserId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "不能删除不是自己的文章"})
			return
		}
		// db.Model(&post).Update(&Post{Title: thisPost.Title, Content: thisPost.Content}).Where("id = ?", thisPost.Id)
		db.Model(&post).Where("id = ?", thisPost.Id).Delete(&post)
		c.JSON(http.StatusOK, gin.H{"msg": "删除文章成功"})
	}
}
