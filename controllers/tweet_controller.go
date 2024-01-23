package controllers

import (
	"net/http"

	"github.com/biqilbooks/dto"
	"github.com/biqilbooks/services"
	"github.com/gin-gonic/gin"
)

func AddTweet(ctx *gin.Context) {
	var body dto.AddTweet

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tweet, err := services.AddTweet(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, tweet)
}

func GetTweetById(ctx *gin.Context) {
	tweetId := ctx.Param("id")

	tweet, err := services.GetTweetById(tweetId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, tweet)

}

func GetAllTweets(ctx *gin.Context) {
	tweets, err := services.GetAllTweets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Respond
	ctx.JSON(http.StatusOK, tweets)
}

func UpdateTweet(ctx *gin.Context) {
	tweetId := ctx.Param("id")
	var body dto.AddTweet

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedTweet, err := services.UpdateTweet(body, tweetId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, updatedTweet)
}


func DeleteTweet(ctx *gin.Context) {

	tweetId := ctx.Param("id")
	tweet, err := services.DeleteTweetbyId(tweetId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Respond
	ctx.JSON(http.StatusOK, tweet)
}
