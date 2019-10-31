package service_test

import (
	"testing"

	"github.com/hernancabral/Twiteer/src/service"

	"github.com/hernancabral/Twiteer/src/domain"

	"github.com/stretchr/testify/assert"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	assert.Equal(t, user, publishedTweet.User)
	assert.Equal(t, text, publishedTweet.Text)
	assert.NotNil(t, publishedTweet.Date)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.EqualError(t, err, "user is required", "Expected error is user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	var text string
	user := "TestUser"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.EqualError(t, err, "text is required", "Expected error is text is required")
}

func TestTweetCanNotHaveMoreThan140Chars(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet

	user := "TestUser"
	text := "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta."

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.EqualError(t, err, "too many characters", "Expected error is too many characters")
}
