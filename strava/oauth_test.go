package strava

import (
	"testing"
)

func TestGetRedirectUrl(t *testing.T) {
	authorizer := NewStravaAuthorizer("some id", "some secret", "read,write", "localhost")
	if authorizer.ClientID != "some id" {
		t.Errorf("Unexpected ID. %s != %s.", authorizer.ClientID, "some id")
	}
	if authorizer.GetRedirectURL() != expectedAuthorizationURL {
		t.Errorf("Unexpectected callback. %s != %s", authorizer.GetRedirectURL(), expectedAuthorizationURL)
	}
}

const expectedAuthorizationURL string = "https://www.strava.com/oauth/authorize?access_type=offline&client_id=some+id&redirect_uri=localhost%2Fcallback&response_type=code&scope=read+write&state=state"
