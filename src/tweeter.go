package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/hernancabral/Twiteer/src/domain"
	"github.com/hernancabral/Twiteer/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")
	tm := service.NewTweetManager()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			tweetUser := c.ReadLine()

			c.Print("Write your tweet: ")

			tweetText := c.ReadLine()

			_, err := tm.PublishTweet(domain.NewTweet(tweetUser, tweetText))

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

			tweet := tm.GetLastTweet()

			c.Println("Usuario: ", tweet.User, "\nTexto: ", tweet.Text, " A las: ", tweet.Date)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows a tweet by Id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the id: ")

			tweetId, _ := strconv.Atoi(c.ReadLine())

			id, _ := tm.GetTweetById(tweetId)

			tweet := tm.GetTweets()[id]

			c.Println("Usuario: ", tweet.User, "\nTexto: ", tweet.Text, " A las: ", tweet.Date)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "howManyTweetsByUser",
		Help: "Shows a tweet by Id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user name: ")

			user := c.ReadLine()

			count := tm.CountTweetsByUser(user)

			c.Println("El usuario: ", user, " twiteo ", count, " veces")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUserTweets",
		Help: "Shows a tweet by Id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the user name: ")

			user := c.ReadLine()

			tweets := tm.GetTweetsByUser(user)

			c.Println("Usuario: ", user, "\n")

			for i := 0; i < len(tweets); i++ {
				c.Println("Texto: ", tweets[i].Text, " A las: ", tweets[i].Date)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tm.GetTweets()

			for i := 0; i < len(tweets); i++ {
				c.Println("Usuario: ", tweets[i].User, "\nTexto: ", tweets[i].Text, " A las: ", tweets[i].Date)
			}

			return
		},
	})

	shell.Run()

}
