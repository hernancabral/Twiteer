package domain

import "time"

type Tweet struct {
	User, Text string
	Date       *time.Time
}

func NewTweet(user, text string) *Tweet {
	t := time.Now()
	return &Tweet{user, text, &t}
}
