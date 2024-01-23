package controllers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetGinRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("tweets", GetAllTweets)
	r.GET("tweets/:id", GetTweetById)
	r.POST("tweets", AddTweet)
	r.PATCH("tweets/:id", UpdateTweet)
	r.DELETE("tweets/:id", DeleteTweet)
}
