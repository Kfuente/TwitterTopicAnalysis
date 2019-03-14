/* Authenticates the twitter login and oauth */

package credentials

import (
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

//LoadCredentials Loads the credentials for twitter API
func LoadCredentials() (client *twittergo.Client) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    APIKey,
		ConsumerSecret: APISecretKey,
	}
	user := oauth1a.NewAuthorizedConfig(AccessToken, AccessTokenSecretKey)

	client = twittergo.NewClient(config, user)
	return
}
