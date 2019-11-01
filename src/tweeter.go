package main

import (
	"github.com/abiosoft/ishell"
	"github.com/hernancabral/Twiteer/src/domain"
	"github.com/hernancabral/Twiteer/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			tweetUser := c.ReadLine()

			c.Print("Write your tweet: ")

			tweetText := c.ReadLine()

			err := service.PublishTweet(domain.NewTweet(tweetUser, tweetText))

			if err != nil {
				c.Print("Error: ", err.Error(), "\n")
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetLastTweet()

			c.Println("Usuario: ", tweet.User, "\nTexto: ", tweet.Text, " A las: ", tweet.Date)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			for i := 0; i < len(tweets); i++ {
				c.Println("Usuario: ", tweets[i].User, "\nTexto: ", tweets[i].Text, " A las: ", tweets[i].Date)
			}

			return
		},
	})

	shell.Run()

}
