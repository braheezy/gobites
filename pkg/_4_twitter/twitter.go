package _4_twitter

/*

Analyze Twitter data.

The complexity of HTTP calls and parsing is handled in client.go. This file is for analysis of that data

*/

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Globally limit historical tweet retrieval
const YearsAgo = 0
const MonthsAgo = 5
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

func ComputeSimilarity(tweets1 []Tweet, tweets2 []Tweet) float64 {
	var document1 []string
	for _, tweet := range tweets1 {
		document1 = append(document1, tweet.Text)
	}
	doc1 := strings.Join(document1, " ")

	var document2 []string
	for _, tweet := range tweets2 {
		document2 = append(document2, tweet.Text)
	}
	doc2 := strings.Join(document2, " ")

	return CosineSimilarity(doc1, doc2)
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
	user1 := "JoeBiden"
	user2 := "elonmusk"

	// Create new TwitterUser for each user
	TwitterUser1 := NewTwitterUser(user1, TwitterClient)

	fmt.Printf("Got a total of %v tweets for %s\n", len(TwitterUser1.Tweets), user1)
	// fmt.Printf("Newest tweet: %s\n", TwitterUser1.Tweets[0])
	// fmt.Printf("Oldest tweet: %s\n", TwitterUser1.Tweets[len(TwitterUser1.Tweets)-1])

	TwitterUser2 := NewTwitterUser(user2, TwitterClient)

	fmt.Printf("Got a total of %v tweets for %s\n", len(TwitterUser2.Tweets), user2)
	// fmt.Printf("Newest tweet: %s\n", TwitterUser2.Tweets[0])
	// fmt.Printf("Oldest tweet: %s\n", TwitterUser2.Tweets[len(TwitterUser2.Tweets)-1])

	fmt.Printf("\nUser %s and %s have a tweet similarity score of: %0.4f\n",
		TwitterUser1.Username,
		TwitterUser2.Username,
		ComputeSimilarity(TwitterUser1.Tweets, TwitterUser2.Tweets),
	)
}
