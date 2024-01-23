package services

import (
	"strconv"

	"github.com/biqilbooks/db"
	"github.com/biqilbooks/dto"
	"github.com/biqilbooks/models"
	"gorm.io/gorm"
)

func AddTweet(body dto.AddTweet) (models.Tweet, error) {
	tweet := models.Tweet{
		Author:  body.Author,
		Content: body.Content,
	}
	err := db.CreateTweet(&tweet)
	if err != nil {
		return tweet, err
	}
	return tweet, nil
}

func GetAllTweets() ([]models.Tweet, error) {
	return db.GetAllTweets()
}

func GetTweetById(tweetId string) (models.Tweet, error) {
	return db.GetTweetById(tweetId)
}

func UpdateTweet(body dto.AddTweet, tweetId string) (models.Tweet, error) {
	tweetIdInt, err := strconv.ParseUint(tweetId, 10, 32)
	if err != nil {
		return models.Tweet{}, err
	}
	tweet := models.Tweet{
		Model: gorm.Model{
			ID: uint(tweetIdInt),
		},
		Author:  body.Author,
		Content: body.Content,
	}
	err = db.UpdateTweet(&tweet)
	if err != nil {
		return models.Tweet{}, err
	}
	return tweet, nil
}

func DeleteTweetbyId(tweetId string) (*models.Tweet, error) {
	return db.DeleteTweet(tweetId)
}
