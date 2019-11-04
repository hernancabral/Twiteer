package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	TextTweet
	URL string
}

func NewImageTweet(user, text, url string) *ImageTweet {
	t := time.Now()
	return &ImageTweet{TextTweet{0, user, text, &t}, url}
}

func (t *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text, t.URL)
}

func (t *ImageTweet) String() string {
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text, t.URL)
}

func (t *ImageTweet) GetURL() string {
	return t.URL
}
