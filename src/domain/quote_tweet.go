package domain

import (
	"fmt"
	"time"
)

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func NewQuoteTweet(user string, text string, tweet Tweet) *QuoteTweet {
	t := time.Now()
	return &QuoteTweet{TextTweet{0, user, text, &t}, tweet}
}

func (t *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"%s\"", t.User, t.Text, t.QuotedTweet)
}

func (t *QuoteTweet) String() string {
	return fmt.Sprintf("@%s: %s \"%s\"", t.User, t.Text, t.QuotedTweet)
}

func (t *QuoteTweet) GetQoutedTweet() Tweet {
	return t.QuotedTweet
}
