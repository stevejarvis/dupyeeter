// Strava uses OAuth2. Details on their scopes, etc. here:
// https://developers.strava.com/docs/authentication/
package strava

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/oauth2"
)

const (
	authorizeURL string = "https://www.strava.com/oauth/authorize"
	tokenURL     string = "https://www.strava.com/oauth/token"
)

type StravaAuthorizer struct {
	*oauth2.Config
}

// Create a new Authorizer, which is the struct on which you can perform the other OAuth2
// requirements for Strava.
func NewStravaAuthorizer(clientId, clientSecret, scopes, serverURL string) *StravaAuthorizer {
	return &StravaAuthorizer{
		Config: &oauth2.Config{
			ClientID:     clientId,
			ClientSecret: clientSecret,
			Scopes:       strings.Split(scopes, ","),
			Endpoint: oauth2.Endpoint{
				AuthURL:  authorizeURL,
				TokenURL: tokenURL,
			},
			RedirectURL: serverURL + "/callback",
		},
	}
}

// Get the URL we need to redirect user to to actually authorize Dupyeeter
// to Strava. Strava's OAuth2 UX.
func (s *StravaAuthorizer) GetRedirectURL() string {
	return s.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

// Once we get a callback from Strava, use the authorization code to get an access token.
// Exchange authorization code for access and refresh token
func (s *StravaAuthorizer) ExchangeCodeForAccessToken(authCode string) (*oauth2.Token, error) {
	// POST https://www.strava.com/oauth/token w/ client ID, secret, & auth code.
	// But hooray, library does this for us.
	token, err := s.Exchange(context.Background(), authCode)
	if err != nil {
		return nil, fmt.Errorf("Failed to get the token. %w", err)
	}

	return token, nil
}
