package service

import (
	"fmt"

	"github.com/hernancabral/Twiteer/src/domain"
)

var Tweet *domain.Tweet
var Tweets []*domain.Tweet
var TweetsByUser map[string][]*domain.Tweet
var Id int

func InitializeService() {
	Tweets = make([]*domain.Tweet, 0)
	TweetsByUser = make(map[string][]*domain.Tweet)
	Id = 1
}

func PublishTweet(tweet *domain.Tweet) (int, error) {
	var err error
	if tweet.User != "" && tweet.Text != "" {
		if len([]rune(tweet.Text)) > 140 {
			return 0, fmt.Errorf("too many characters")
		}
		tweet.Id = Id
		Id++
		Tweets = append(Tweets, tweet)
		TweetsByUser[tweet.User] = append(TweetsByUser[tweet.User], tweet)
		return tweet.Id, err
	} else {
		if tweet.User == "" {
			return 0, fmt.Errorf("user is required")
		} else {
			return 0, fmt.Errorf("text is required")
		}
	}
}

func GetTweets() []*domain.Tweet {
	return Tweets
}

func GetLastTweet() *domain.Tweet {
	return GetTweets()[len(GetTweets())-1]
}

func GetTweetById(IdToSearch int) (int, error) {
	var err error
	for i := range Tweets {
		if Tweets[i].Id == IdToSearch {
			return i, err
		}
	}
	return 0, fmt.Errorf("Tweet not found")
}

func CountTweetsByUser(user string) int {
	var cont int = 0
	for i := range Tweets {
		if Tweets[i].User == user {
			cont++
		}
	}
	return cont
}

func GetTweetsByUser(user string) []*domain.Tweet {
	return TweetsByUser[user]
}
