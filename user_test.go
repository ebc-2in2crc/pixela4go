package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateUserCreateRequestParameter(t *testing.T) {
	client := New("name", "token")
	input := &UserCreateInput{
		AgreeTermsOfService: Bool(true),
		NotMinor:            Bool(true),
		ThanksCode:          String("thanks-code"),
	}
	param, err := client.User().createCreateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL + "/users")
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	s := `{"token":"token","username":"name","AgreeTermsOfService":"yes","NotMinor":"yes","thanksCode":"thanks-code"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestUserCreate(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &UserCreateInput{
		AgreeTermsOfService: Bool(true),
		NotMinor:            Bool(true),
		ThanksCode:          String("thanks-code"),
	}
	result, err := client.User().Create(input)

	testSuccess(t, result, err)
}

func TestUserCreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &UserCreateInput{
		AgreeTermsOfService: Bool(true),
		NotMinor:            Bool(true),
		ThanksCode:          String("thanks-code"),
	}
	result, err := client.User().Create(input)

	testAPIFailedResult(t, result, err)
}

func TestUserCreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &UserCreateInput{
		AgreeTermsOfService: Bool(true),
		NotMinor:            Bool(true),
		ThanksCode:          String("thanks-code"),
	}
	_, err := client.User().Create(input)

	testPageNotFoundError(t, err)
}

func TestCreateUserUpdateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &UserUpdateInput{
		NewToken:   String("newtoken"),
		ThanksCode: String("thanks-code"),
	}
	param, err := client.User().createUpdateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"newToken":"newtoken","thanksCode":"thanks-code"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestUserUpdate(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &UserUpdateInput{
		NewToken:   String("newToken"),
		ThanksCode: String("thanks-code"),
	}
	result, err := client.User().Update(input)

	testSuccess(t, result, err)
}

func TestUserUpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &UserUpdateInput{
		NewToken:   String("newToken"),
		ThanksCode: String("thanks-code"),
	}
	result, err := client.User().Update(input)

	testAPIFailedResult(t, result, err)

	if client.Token != token {
		t.Errorf("got: %s\nwant: %s", client.Token, token)
	}
}

func TestUserUpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &UserUpdateInput{
		NewToken:   String("newToken"),
		ThanksCode: String("thanks-code"),
	}
	_, err := client.User().Update(input)

	testPageNotFoundError(t, err)

	if client.Token != token {
		t.Errorf("got: %s\nwant: %s", client.Token, token)
	}
}

func TestCreateUserDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	param, err := client.User().createDeleteRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestUserDelete(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	result, err := client.User().Delete()

	testSuccess(t, result, err)
}

func TestUserDeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	result, err := client.User().Delete()

	testAPIFailedResult(t, result, err)
}

func TestUserDeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	_, err := client.User().Delete()

	testPageNotFoundError(t, err)
}
