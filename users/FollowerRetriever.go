package users

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/kurrik/twittergo"
)

const (
	count                    = "100"
	friendsTwitterPartialURL = "/1.1/friends/list.json?%v"
)

/*
RetrieveFollowersForASpecificUser returns the list of people that a users
follows
*/
func RetrieveFollowersForASpecificUser(user string, client *twittergo.Client) []string {
	var (
		query       url.Values
		req         *http.Request
		apiResponse *twittergo.APIResponse
		err         error
	)

	query = url.Values{}
	query.Set("screen_name", user)
	query.Set("count", count)
	endpoint := fmt.Sprintf(friendsTwitterPartialURL, query.Encode())

	if req, err = http.NewRequest("GET", endpoint, nil); err != nil {
		fmt.Printf("Could not parse request for user follower list: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Accept-Encoding", "gzip, deflate")

	if apiResponse, err = client.SendRequest(req); err != nil {
		fmt.Printf("Could not send request for user follower list: %v\n", err)
		os.Exit(1)
	}

	type UserLists struct {
		CustomID []twittergo.User `json:"users"`
	}

	userList := &UserLists{}

	if err = apiResponse.Parse(userList); err != nil {
		fmt.Printf("Could not parse the follower list response: %v\n", err)
		os.Exit(1)
	}

	followerScreenNames := make([]string, 0)
	for _, list := range (*userList).CustomID {
		followerScreenNames = append(followerScreenNames, getUserName(list))
	}

	return followerScreenNames
}

func getUserName(user twittergo.User) (followerScreenName string) {
	followerScreenName = user.ScreenName()
	return
}
