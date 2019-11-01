package service_test

import (
	"testing"

	"github.com/hernancabral/Twiteer/src/service"

	"github.com/hernancabral/Twiteer/src/domain"

	"github.com/stretchr/testify/assert"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	service.InitializeService()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetLastTweet()
	assert.Equal(t, user, publishedTweet.User)
	assert.Equal(t, text, publishedTweet.Text)
	assert.NotNil(t, publishedTweet.Date)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	service.InitializeService()
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
	service.InitializeService()
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
	service.InitializeService()
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

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data
	tweet = domain.NewTweet("Usuario", "texto1")
	secondTweet = domain.NewTweet("Usuario", "texto2")

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	assert.EqualValues(t, 2, len(publishedTweets))

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	assert.EqualValues(t, "Usuario", firstPublishedTweet.User)
	assert.EqualValues(t, "texto1", firstPublishedTweet.Text)
	assert.EqualValues(t, "Usuario", secondPublishedTweet.User)
	assert.EqualValues(t, "texto2", secondPublishedTweet.Text)
}
