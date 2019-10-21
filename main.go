package main

import (
	"TwitterTopicAnalysis/credentials"
	"TwitterTopicAnalysis/timeline"
	"TwitterTopicAnalysis/users"
	"fmt"
	"net/http"
	"os"

	"github.com/kurrik/twittergo"
)

func main() {
	var (
		err    error
		client *twittergo.Client
		req    *http.Request
		resp   *twittergo.APIResponse
		user   *twittergo.User
	)

	client = credentials.LoadCredentials()

	req, err = http.NewRequest("GET", "/1.1/account/verify_credentials.json", nil)
	if err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}
	resp, err = client.SendRequest(req)
	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}
	user = &twittergo.User{}
	err = resp.Parse(user)
	if err != nil {
		fmt.Printf("Problem parsing response: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("ID:                   %v\n", user.Id())
	fmt.Printf("Name:                 %v\n", user.Name())
	if resp.HasRateLimit() {
		fmt.Printf("Rate limit:           %v\n", resp.RateLimit())
		fmt.Printf("Rate limit remaining: %v\n", resp.RateLimitRemaining())
		fmt.Printf("Rate limit reset:     %v\n", resp.RateLimitReset())
	} else {
		fmt.Printf("Could not parse rate limit from response.\n")
	}

	currentTimeline := timeline.RetrieveUserTimeline("FLoloz", client)

	for _, tweet := range *currentTimeline {
		fmt.Printf("%+v\n", tweet.Text())
	}

	followerScreenNameList := users.RetrieveFollowersForASpecificUser("FLoloz", client)

	for _, followerScreenName := range followerScreenNameList {
		fmt.Printf("User: %v\n", followerScreenName)
	}

}
