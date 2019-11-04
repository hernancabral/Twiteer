package service_test

import (
	"testing"

	"github.com/hernancabral/Twiteer/src/service"

	"github.com/hernancabral/Twiteer/src/domain"

	"github.com/stretchr/testify/assert"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tm := service.NewTweetManager()
	var tweet *domain.TextTweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	tm.PublishTweet(tweet)

	// Validation
	publishedTweet := tm.GetLastTweet()
	assert.Equal(t, user, publishedTweet.GetUser())
	assert.Equal(t, text, publishedTweet.GetText())
	assert.NotNil(t, publishedTweet.GetDate())
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tm := service.NewTweetManager()
	var tweet *domain.TextTweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.EqualError(t, err, "user is required", "Expected error is user is required")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tm := service.NewTweetManager()
	var tweet *domain.TextTweet

	var text string
	user := "TestUser"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.EqualError(t, err, "text is required", "Expected error is text is required")
}

func TestTweetCanNotHaveMoreThan140Chars(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet *domain.TextTweet

	user := "TestUser"
	text := "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta."

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tm.PublishTweet(tweet)

	// Validation
	assert.EqualError(t, err, "too many characters", "Expected error is too many characters")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet, secondTweet *domain.TextTweet // Fill the tweets with data
	tweet = domain.NewTextTweet("Usuario", "texto1")
	secondTweet = domain.NewTextTweet("Usuario", "texto2")

	// Operation
	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tm.GetTweets()
	assert.EqualValues(t, 2, len(publishedTweets))

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	assert.EqualValues(t, "Usuario", firstPublishedTweet.GetUser())
	assert.EqualValues(t, "texto1", firstPublishedTweet.GetText())
	assert.EqualValues(t, "Usuario", secondPublishedTweet.GetUser())
	assert.EqualValues(t, "texto2", secondPublishedTweet.GetText())
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tm := service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tm.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := tm.GetTweetById(id)

	//isValidTweet(t, publishedTweet, id, user, text)

	assert.Equal(t, id, tm.GetTweets()[publishedTweet].GetId())
	assert.Equal(t, user, tm.GetTweets()[publishedTweet].GetUser())
	assert.Equal(t, text, tm.GetTweets()[publishedTweet].GetText())

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)
	// Operation
	count := tm.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tm := service.NewTweetManager()
	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
	// publish the 3 tweets
	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)

	// Operation
	tweets := tm.GetTweetsByUser(user)

	// Validation
	assert.Equal(t, 2, len(tweets))

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	assert.Equal(t, user, firstPublishedTweet.GetUser())
	assert.Equal(t, text, firstPublishedTweet.GetText())

	assert.Equal(t, user, secondPublishedTweet.GetUser())
	assert.Equal(t, secondText, secondPublishedTweet.GetText())

}

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	assert.Equal(t, expectedText, text)
}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	assert.Equal(t, expectedText, text)

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	assert.Equal(t, expectedText, text)

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	assert.Equal(t, expectedText, text)
}
