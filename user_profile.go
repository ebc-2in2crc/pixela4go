package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

)

// A UserProfile manages communication with the Pixela user profile API.
type UserProfile struct {
	UserName   string
	Token      string
	httpClient HTTPClient
}

// Update updates the profile information for the user corresponding to username.
func (u *UserProfile) Update(input *UserProfileUpdateInput) (*Result, error) {
	return u.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext updates the profile information for the user corresponding to username.
func (u *UserProfile) UpdateWithContext(ctx context.Context, input *UserProfileUpdateInput) (*Result, error) {
	param, err := u.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, fmt.Errorf("failed to create user profile update parameter: %w", err)
	}

	return doRequestAndParseResponse(ctx, u.httpClient, param)
}

// UserProfileUpdateInput is input of UserProfile.Update().
type UserProfileUpdateInput struct {
	DisplayName       *string  `json:"displayName,omitempty"`
	GravatarIconEmail *string  `json:"gravatarIconEmail,omitempty"`
	Title             *string  `json:"title,omitempty"`
	Timezone          *string  `json:"timezone,omitempty"`
	AboutURL          *string  `json:"aboutURL,omitempty"`
	ContributeURLs    []string `json:"contributeURLs,omitempty"`
	PinnedGraphID     *string  `json:"pinnedGraphID,omitempty"`
}

func (u *UserProfile) createUpdateRequestParameter(input *UserProfileUpdateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/@%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   b,
	}, nil
}

// URL outputs the profile of the user specified by username in html format.
func (u *UserProfile) URL() string {
	return fmt.Sprintf(APIBaseURL+"/@%s", u.UserName)
}
