package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateUserProfileUpdateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &UserProfileUpdateInput{
		DisplayName:       String("displayName"),
		GravatarIconEmail: String("gravatarIconEmail"),
		Title:             String("title"),
		Timezone:          String("timezone"),
		AboutURL:          String("aboutURL"),
		ContributeURLs:    []string{"hoge.com"},
		PinnedGraphID:     String("pinnedGraphID"),
	}
	param, err := client.UserProfile().createUpdateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/@%s", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"displayName":"displayName","gravatarIconEmail":"gravatarIconEmail","title":"title","timezone":"timezone","aboutURL":"aboutURL","contributeURLs":["hoge.com"],"pinnedGraphID":"pinnedGraphID"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestUserProfileUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &UserProfileUpdateInput{
		DisplayName:       String("displayName"),
		GravatarIconEmail: String("gravatarIconEmail"),
		Title:             String("title"),
		Timezone:          String("timezone"),
		AboutURL:          String("aboutURL"),
		ContributeURLs:    []string{"hoge.com"},
		PinnedGraphID:     String("pinnedGraphID"),
	}
	result, err := client.UserProfile().Update(input)

	testSuccess(t, result, err)
}

func TestUserProfileUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &UserProfileUpdateInput{
		DisplayName:       String("displayName"),
		GravatarIconEmail: String("gravatarIconEmail"),
		Title:             String("title"),
		Timezone:          String("timezone"),
		AboutURL:          String("aboutURL"),
		ContributeURLs:    []string{"hoge.com"},
		PinnedGraphID:     String("pinnedGraphID"),
	}
	result, err := client.UserProfile().Update(input)

	testAPIFailedResult(t, result, err)
}

func TestUserProfileUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &UserProfileUpdateInput{
		DisplayName:       String("displayName"),
		GravatarIconEmail: String("gravatarIconEmail"),
		Title:             String("title"),
		Timezone:          String("timezone"),
		AboutURL:          String("aboutURL"),
		ContributeURLs:    []string{"hoge.com"},
		PinnedGraphID:     String("pinnedGraphID"),
	}
	_, err := client.UserProfile().Update(input)

	testPageNotFoundError(t, err)
}

func TestUserProfile_URL(t *testing.T) {
	client := New(userName, token)
	url := client.UserProfile().URL()
	expect := fmt.Sprintf(APIBaseURL+"/@%s", userName)
	if url != expect {
		t.Errorf("got: %s\nwant: %s", url, expect)
	}
}
