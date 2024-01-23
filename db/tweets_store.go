package db

import (
	"github.com/biqilbooks/models"
)

func CreateTweet(tweet *models.Tweet) error {
	return DB.Create(&tweet).Error
}

func UpdateTweet(tweet *models.Tweet) error {
	err := DB.Save(*tweet).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTweet(tweetId string) (*models.Tweet, error) {
	var deletedTweet models.Tweet

	// Find the tweet by its ID
	result := DB.First(&deletedTweet, tweetId)
	if result.Error != nil {
		return nil, result.Error
	}

	// Delete the tweet
	result = DB.Delete(&deletedTweet)
	if result.Error != nil {
		return nil, result.Error
	}

	return &deletedTweet, nil
}

func GetAllTweets() ([]models.Tweet, error) {
	var tweets []models.Tweet
	err := DB.Find(&tweets).Error
	if err != nil {
		return tweets, err
	}
	return tweets, nil
}

func GetTweetById(tweetId string) (models.Tweet, error) {
	var tweet models.Tweet
	err := DB.First(&tweet, "id=?", tweetId).Error
	if err != nil {
		return tweet, err
	}
	return tweet, nil
}
