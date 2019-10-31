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

			tweet := service.GetTweet()

			c.Println("Usuario: ", tweet.User, "\nTexto: ", tweet.Text, " A las: ", tweet.Date)

			return
		},
	})

	shell.Run()

}
