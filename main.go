package main

import (
	"fmt"
	"github.com/hjr265/too"
	"log"
	"net"
)

type Like struct {
	user too.User
	show too.Item
}

type Client struct {
	// Public access
	Like       chan Like
	Suggestion chan too.User
	Close      chan bool
	// Private access
	connection *too.Engine
}

const NUM_RECOMMENDATIONS = 3

var redisConnection *too.Engine = nil

func main() {

	c := NewClient()

	// Add sample data in the database
	addSampleData(c.Like)

	// Get suggestions:
	fmt.Println("Suggestions:")
	getSuggestions(c.Suggestion, "Albert")
}

func NewClient() Client {

	redisAddr, _ := net.ResolveTCPAddr("tcp", ":6379")
	conn, err := too.New(redisAddr, "tvshows")

	if err != nil {
		log.Fatal(err)
	}

	client := Client{make(chan Like), make(chan too.User), make(chan bool), conn}
	go client.run()
	return client
}

func (c *Client) run() {

	for {
		select {
		case like := <-c.Like:
			fmt.Println("Adding recommendation", like.user, like.show)
			c.connection.Likes.Add(like.user, like.show)

		case user := <-c.Suggestion:
			fmt.Println("Suggesting")
			items, _ := c.connection.Suggestions.For(user, NUM_RECOMMENDATIONS)
			for _, item := range items {
				fmt.Println(item)
			}

		case <-c.Close:
			return

		}

	}
}

func getSuggestions(c chan too.User, user too.User) {
	c <- user
}

func addSampleData(c chan Like) {
	likes := []Like{
		{"Albert", "Game of Thrones"},
		{"Albert", "The Big Bang Theory"},
		{"Albert", "Fargo"},
		{"Albert", "How I Met Your Mother"},
		{"Albert", "Breaking Bad"},
		{"Albert", "The Strain"},
		{"Albert", "Six Feet Under"},

		{"Noemí", "Grey's Anatomy"},
		{"Noemí", "The Good Wife"},
		{"Noemí", "Game of Thrones"},
		{"Noemí", "The Strain"},

		{"Pepe", "Grey's Anatomy"},
		{"Pepe", "Six Feet Under"},
		{"Pepe", "How I Met Your Mother"},

		{"Paco", "The Good Wife"},
		{"Paco", "Game of Thrones"},
		{"Paco", "Fargo"},
	}

	for _, like := range likes {
		c <- like
	}
}
