package service

var Tweet string

func PublishTweet(tweet string) {
	Tweet = "This is my first tweet!"
}

func GetTweet() string {
	return Tweet
}
