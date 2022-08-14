package _4_twitter

/*
This file contains all the functionality to interact with the API.

`TwitterClient` wraps http.Client and adds Twitter API info.

Helper functions `Get` and `ParseResponse` do the work of making generic HTTP calls.

The rest of the file contains specific API calls the client can make.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A struct to hold info for a Client entity that can call Twitter API
type TwitterClient struct {
	// The Bearer Token to use
	Token string
	// The URL for the API
	URL string
	// native http client
	httpClient http.Client
}

// Create a new TwitterClient type
func NewTwitterClient(token string) TwitterClient {
	return TwitterClient{token, "https://api.twitter.com/2", http.Client{}}
}

// Perform a request using TwitterClient
func (c *TwitterClient) Get(url string) []byte {
	// Build up a request so we can attach a header
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	rawBody := ParseResponse(c.httpClient.Do(req))

	return rawBody
}

func ParseResponse(response *http.Response, err error) []byte {
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return body
}

/*
	API caller functions.

	A type is defined per API call to handle that call's unique data response.
	Each function typically:
		- Forms the URL for the API endpoint
		- Calls the API
		- Handles the JSON response, passing through a struct cause that seems to be the approach in Go.
		- Returns easy to use data (not the whole struct)
*/

//--------------------------------------------------------------
// 	UserInfo: Get the basic facts about a user entity
//--------------------------------------------------------------
type UserInfo struct {
	Info []struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

// Given a username, get all the other user info stuff
func (c TwitterClient) GetUserInfo(username string) (string, string, string) {
	url := fmt.Sprintf("%s/users/by?usernames=%s", c.URL, username)

	rawBody := c.Get(url)

	var u UserInfo
	json.Unmarshal(rawBody, &u)

	return u.Info[0].Id, u.Info[0].Name, u.Info[0].Username
}

//--------------------------------------------------------------
// 	UserTweets: Collect a bunch of a user's tweets
//--------------------------------------------------------------
type UserTweets struct {
	Tweets []Tweet `json:"data"`
	// Pagination support
	Meta struct {
		NextToken   string `json:"next_token"`
		ResultCount int    `json:"result_count"`
		NewestId    string `json:"newest_id"`
		OldestId    string `json:"oldest_id"`
	} `json:"meta"`
}

type Tweet struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func (c TwitterClient) GetUserTweets(userId string) (TweetList []Tweet) {
	// Define the initial URL
	baseUrl := fmt.Sprintf("%s/users/%s/tweets?max_results=%v&start_time=%v", c.URL, userId, MaxResults, TweetLimit)

	// Make the first call like normal
	rawBody := c.Get(baseUrl)
	// Parse the response
	var currentTweets UserTweets
	json.Unmarshal(rawBody, &currentTweets)

	// Create a separate list to hold all tweets obtained
	TweetList = append(TweetList, currentTweets.Tweets...)

	// Paginate until the result count drops, indicating some stop condition was hit
	for currentTweets.Meta.ResultCount == MaxResults {
		// Make new url
		url := fmt.Sprintf("%s&pagination_token=%s", baseUrl, currentTweets.Meta.NextToken)
		// Get next page of results
		rawBody = c.Get(url)
		json.Unmarshal(rawBody, &currentTweets)
		TweetList = append(TweetList, currentTweets.Tweets...)
	}

	return TweetList
}

//--------------------------------------------------------------
