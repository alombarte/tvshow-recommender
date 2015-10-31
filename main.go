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

const NUM_RECOMMENDATIONS = 3

var redisConnection *too.Engine = nil

func New() *too.Engine {
	if redisConnection == nil {
		redisAddr, _ := net.ResolveTCPAddr("tcp", ":6379")
		conn, err := too.New(redisAddr, "movies")
		if err != nil {
			log.Fatal(err)
		}
		return conn
	}

	return redisConnection
}

func main() {

	// Add sample data in the database
	addSampleData()

	// Print recommendations for user
	connection := New()
	items, _ := connection.Suggestions.For("Albert", NUM_RECOMMENDATIONS)
	for _, item := range items {
		fmt.Println(item)
	}
}

func addSampleData() {
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
		{"Noemí", "The Prestige"},
		{"Noemí", "The Matrix"},
		{"Noemí", "The Strain"},

		{"Pepe", "Grey's Anatomy"},
		{"Pepe", "Six Feet Under"},
		{"Pepe", "How I Met Your Mother"},

		{"Paco", "The Good Wife"},
		{"Paco", "Game of Thrones"},
		{"Paco", "Fargo"},
	}

	connection := New()
	for _, like := range likes {
		connection.Likes.Add(like.user, like.show)
	}
}

func insertLike(user, show string) {

}
