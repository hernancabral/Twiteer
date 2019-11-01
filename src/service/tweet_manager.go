package service

import (
	"fmt"

	"github.com/hernancabral/Twiteer/src/domain"
)

var Tweet *domain.Tweet
var Tweets []*domain.Tweet

func InitializeService() {
	Tweets = make([]*domain.Tweet, 0)
}

func PublishTweet(tweet *domain.Tweet) error {
	var err error
	if tweet.User != "" && tweet.Text != "" {
		if len([]rune(tweet.Text)) > 140 {
			return fmt.Errorf("too many characters")
		}
		Tweets = append(Tweets, tweet)
		return err
	} else {
		if tweet.User == "" {
			return fmt.Errorf("user is required")
		} else {
			return fmt.Errorf("text is required")
		}
	}
}

func GetTweets() []*domain.Tweet {
	return Tweets
}

func GetLastTweet() *domain.Tweet {
	return GetTweets()[len(GetTweets())-1]
}
