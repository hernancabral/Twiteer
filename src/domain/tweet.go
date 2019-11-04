package domain

import "time"

type Tweet interface {
	PrintableTweet() string
	String() string
	GetUser() string
	GetText() string
	GetId() int
	SetId(id int)
	GetDate() *time.Time
}
