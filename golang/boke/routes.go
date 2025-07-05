package main

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB, authMiddleware gin.HandlerFunc) {
	// 公共路由
	public := r.Group("/api")

	// 用户认证
	public.POST("/login", loginHandler(db))
	public.POST("/register", registerHandler(db))

	//private 这个路由用了登录token验证
	private := r.Group("/api").Use(authMiddleware)
	// 文章
	private.POST("/createPost", createPostHandler(db))
	private.GET("/listPost", listPostHandler(db))
	private.POST("/updatePost", updatePostHandler(db))
	private.POST("/deletePost", deletePostHandler(db))

	// 文章评论
	private.POST("/createComment", createCommentHandler(db))
	private.GET("/listComment", listCommentHandler(db))

}
