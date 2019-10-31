package service

import (
	"fmt"

	"github.com/hernancabral/Twiteer/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	var err error
	if tweet.User != "" && tweet.Text != "" {
		if len([]rune(tweet.Text)) > 140 {
			return fmt.Errorf("too many characters")
		}
		Tweet = tweet
		return err
	} else {
		if tweet.User == "" {
			return fmt.Errorf("user is required")
		} else {
			return fmt.Errorf("text is required")
		}
	}
}

func GetTweet() *domain.Tweet {
	return Tweet
}
