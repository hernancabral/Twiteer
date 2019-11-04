package service

import (
	"fmt"

	"github.com/hernancabral/Twiteer/src/domain"
)

type TweetManager struct {
	Tweets       []domain.Tweet
	TweetsByUser map[string][]domain.Tweet
	Id           int
}

func NewTweetManager() *TweetManager {
	tm := TweetManager{}
	tm.Tweets = make([]domain.Tweet, 0)
	tm.TweetsByUser = make(map[string][]domain.Tweet)
	return &tm
}

func (tm *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {
	var err error
	if tweet.GetUser() != "" && tweet.GetText() != "" {
		if len([]rune(tweet.GetText())) > 140 {
			return 0, fmt.Errorf("too many characters")
		}
		tweet.SetId(tm.Id)
		tm.Id++
		tm.Tweets = append(tm.Tweets, tweet)
		tm.TweetsByUser[tweet.GetUser()] = append(tm.TweetsByUser[tweet.GetUser()], tweet)
		return tweet.GetId(), err
	} else {
		if tweet.GetUser() == "" {
			return 0, fmt.Errorf("user is required")
		} else {
			return 0, fmt.Errorf("text is required")
		}
	}
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	return tm.Tweets
}

func (tm *TweetManager) GetLastTweet() domain.Tweet {
	return tm.GetTweets()[len(tm.GetTweets())-1]
}

func (tm *TweetManager) GetTweetById(IdToSearch int) (int, error) {
	var err error
	for i := range tm.Tweets {
		if tm.Tweets[i].GetId() == IdToSearch {
			return i, err
		}
	}
	return 0, fmt.Errorf("Tweet not found")
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	var cont int = 0
	for i := range tm.Tweets {
		if tm.Tweets[i].GetUser() == user {
			cont++
		}
	}
	return cont
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tm.TweetsByUser[user]
}
