package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	Id         int
	User, Text string
	Date       *time.Time
}

func NewTextTweet(user, text string) *TextTweet {
	t := time.Now()
	return &TextTweet{0, user, text, &t}
}

func (t *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *TextTweet) String() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *TextTweet) GetUser() string {
	return t.User
}

func (t *TextTweet) GetText() string {
	return t.Text
}

func (t *TextTweet) GetId() int {
	return t.Id
}

func (t *TextTweet) SetId(id int) {
	t.Id = id
}

func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}
