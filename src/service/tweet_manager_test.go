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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := service.GetTweetById(id)

	//isValidTweet(t, publishedTweet, id, user, text)

	assert.Equal(t, id, service.GetTweets()[publishedTweet].Id)
	assert.Equal(t, user, service.GetTweets()[publishedTweet].User)
	assert.Equal(t, text, service.GetTweets()[publishedTweet].Text)

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	assert.Equal(t, 2, len(tweets))

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	assert.Equal(t, user, firstPublishedTweet.User)
	assert.Equal(t, text, firstPublishedTweet.Text)

	assert.Equal(t, user, secondPublishedTweet.User)
	assert.Equal(t, secondText, secondPublishedTweet.Text)

}
