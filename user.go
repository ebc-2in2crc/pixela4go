package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

)

// A User manages communication with the Pixela user API.
type User struct {
	UserName string
	Token    string
}

// Create creates a new Pixela user.
func (u *User) Create(input *UserCreateInput) (*Result, error) {
	return u.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new Pixela user.
func (u *User) CreateWithContext(ctx context.Context, input *UserCreateInput) (*Result, error) {
	param, err := u.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, fmt.Errorf("failed to create user create parameter: %w", err)
	}

	return doRequestAndParseResponse(ctx, param)
}

// UserCreateInput is input of User.Create().
type UserCreateInput struct {
	// AgreeTermsOfService is a required field
	AgreeTermsOfService *bool
	// NotMinor is a required field
	NotMinor   *bool
	ThanksCode *string
}

func (u *User) createCreateRequestParameter(input *UserCreateInput) (*requestParameter, error) {
	b, err := json.Marshal(struct {
		Token               string `json:"token"`
		UserName            string `json:"username"`
		AgreeTermsOfService string `json:"AgreeTermsOfService"`
		NotMinor            string `json:"NotMinor"`
		ThanksCode          string `json:"thanksCode,omitempty"`
	}{
		Token:               u.Token,
		UserName:            u.UserName,
		AgreeTermsOfService: boolToString(BoolValue(input.AgreeTermsOfService)),
		NotMinor:            boolToString(BoolValue(input.NotMinor)),
		ThanksCode:          StringValue(input.ThanksCode),
	})
	if err != nil {
		return &requestParameter{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    APIBaseURLForV1 + "/users",
		Header: map[string]string{},
		Body:   b,
	}, nil
}

func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// Update updates the authentication token for the specified user.
func (u *User) Update(input *UserUpdateInput) (*Result, error) {
	return u.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext updates the authentication token for the specified user.
func (u *User) UpdateWithContext(ctx context.Context, input *UserUpdateInput) (*Result, error) {
	param, err := u.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, fmt.Errorf("failed to create user update parameter: %w", err)
	}

	return doRequestAndParseResponse(ctx, param)
}

// UserUpdateInput is input of User.Update().
type UserUpdateInput struct {
	// NewToken is a required field
	NewToken          *string
	ThanksCode        *string
	AllowAIProcessing *bool
}

func (u *User) createUpdateRequestParameter(input *UserUpdateInput) (*requestParameter, error) {
	b, err := json.Marshal(struct {
		NewToken          string `json:"newToken"`
		ThanksCode        string `json:"thanksCode,omitempty"`
		AllowAIProcessing *bool  `json:"allowAIProcessing,omitempty"`
	}{
		NewToken:          StringValue(input.NewToken),
		ThanksCode:        StringValue(input.ThanksCode),
		AllowAIProcessing: input.AllowAIProcessing,
	})
	if err != nil {
		return &requestParameter{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   b,
	}, nil
}

// Delete deletes the specified registered user.
func (u *User) Delete() (*Result, error) {
	return u.DeleteWithContext(context.Background())
}

// DeleteWithContext deletes the specified registered user.
func (u *User) DeleteWithContext(ctx context.Context) (*Result, error) {
	return doRequestAndParseResponse(ctx, u.createDeleteRequestParameter())
}

func (u *User) createDeleteRequestParameter() *requestParameter {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   []byte{},
	}
}
