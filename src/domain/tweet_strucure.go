package domain

import (
	"time"
)

type Tweet struct {
	Id         int
	User, Text string
	Date       *time.Time
}

func NewTweet(user, text string) *Tweet {
	t := time.Now()
	return &Tweet{0, user, text, &t}
}
