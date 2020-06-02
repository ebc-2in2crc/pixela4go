package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A User manages communication with the Pixela user API.
type User struct {
	UserName string
	Token    string
}

// Create creates a new Pixela user.
func (u *User) Create(input *UserCreateInput) (*Result, error) {
	param, err := u.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create user create parameter")
	}

	return doRequestAndParseResponse(param)
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
	create := &userCreate{
		Token:               u.Token,
		UserName:            u.UserName,
		AgreeTermsOfService: boolToString(BoolValue(input.AgreeTermsOfService)),
		NotMinor:            boolToString(BoolValue(input.NotMinor)),
		ThanksCode:          StringValue(input.ThanksCode),
	}
	b, err := json.Marshal(create)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    APIBaseURL + "/users",
		Header: map[string]string{},
		Body:   b,
	}, nil
}

type userCreate struct {
	Token               string `json:"token"`
	UserName            string `json:"username"`
	AgreeTermsOfService string `json:"AgreeTermsOfService"`
	NotMinor            string `json:"NotMinor"`
	ThanksCode          string `json:"thanksCode,omitempty"`
}

func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// Update updates the authentication token for the specified user.
func (u *User) Update(input *UserUpdateInput) (*Result, error) {
	param, err := u.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create user update parameter")
	}

	return doRequestAndParseResponse(param)
}

// UserUpdateInput is input of User.Update().
type UserUpdateInput struct {
	// NewToken is a required field
	NewToken   *string
	ThanksCode *string
}

func (u *User) createUpdateRequestParameter(input *UserUpdateInput) (*requestParameter, error) {
	update := userUpdate{
		NewToken:   StringValue(input.NewToken),
		ThanksCode: StringValue(input.ThanksCode),
	}
	b, err := json.Marshal(update)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   b,
	}, nil
}

type userUpdate struct {
	NewToken   string `json:"newToken"`
	ThanksCode string `json:"thanksCode,omitempty"`
}

// Delete deletes the specified registered user.
func (u *User) Delete() (*Result, error) {
	param, err := u.createDeleteRequestParameter()
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create user delete parameter")
	}

	return doRequestAndParseResponse(param)
}

func (u *User) createDeleteRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s", u.UserName),
		Header: map[string]string{userToken: u.Token},
		Body:   []byte{},
	}, nil
}
