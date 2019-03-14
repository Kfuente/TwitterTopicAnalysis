package timeline

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/kurrik/twittergo"
)

const (
	count   int = 100
	urltmpl     = "/1.1/statuses/user_timeline.json?%v"
)

func RetrieveUserTimeline(username string, client *twittergo.Client) *twittergo.Timeline {
	var (
		err         error
		req         *http.Request
		query       url.Values
		apiResponse *twittergo.APIResponse
	)

	query = url.Values{}
	query.Set("count", fmt.Sprintf("%v", count))
	query.Set("screen_name", username)
	endpoint := fmt.Sprintf(urltmpl, query.Encode())

	if req, err = http.NewRequest("GET", endpoint, nil); err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
		os.Exit(1)
	}

	if apiResponse, err = client.SendRequest(req); err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		os.Exit(1)
	}

	currentTimeline := &twittergo.Timeline{}
	apiResponse.Parse(currentTimeline)

	return currentTimeline
}
