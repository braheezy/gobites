package _4_twitter

/*

Analyze Twitter data.

The complexity of HTTP calls and parsing is handled in client.go. This file is for analysis of that data

*/

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Globally limit historical tweet retrieval
const YearsAgo = 0
const MonthsAgo = 1
const DaysAgo = 0

// Limit for any pagination operations
const MaxResults = 100

// Force a limit to how long ago tweets should be obtained. People tweet a lot!
// Twitter API wants YYYY-MM-DDTHH:mm:ssZ a.k.a RFC3339 format
var TweetLimit = time.Now().AddDate(-YearsAgo, -MonthsAgo, -DaysAgo).Format(time.RFC3339)

// Custom type to hold a bunch of Twitter info for a user
type TwitterUser struct {
	// Various ways to refer to a user
	Id       string
	Name     string
	Username string
	// Collection of tweets they user has posted
	Tweets []Tweet
}

func NewTwitterUser(handle string, c TwitterClient) TwitterUser {
	id, name, username := c.GetUserInfo(handle)
	tweets := c.GetUserTweets(id)

	return TwitterUser{id, name, username, tweets}
}

func Run() {
	// Get secrets from env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Client that can make Twitter API calls.
	TwitterClient := NewTwitterClient(os.Getenv("TWITTER_BEARER_TOKEN"))

	// Define user handles to inspect
	// TODO: Let user input handles to inspect
	user1 := "elonmusk"

	// Create new TwitterUser for each user
	TwitterUser1 := NewTwitterUser(user1, TwitterClient)

	fmt.Printf("Got a total of %v tweets!\n", len(TwitterUser1.Tweets))
	fmt.Printf("Newest tweet: %s\n", TwitterUser1.Tweets[0])
	fmt.Printf("Oldest tweet: %s\n", TwitterUser1.Tweets[len(TwitterUser1.Tweets)-1])
}
