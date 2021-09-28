package pixela

import (
	"net/http"
	"os"
	"testing"
)

func testE2EUserCreate(t *testing.T) {
	input := &UserCreateInput{
		AgreeTermsOfService: Bool(true),
		NotMinor:            Bool(true),
	}
	result, err := e2eClient.User().Create(input)
	if err != nil {
		t.Errorf("User.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("User.Create() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("User.Create() got: %+v\nwant: %d", result, http.StatusOK)
	}
}

func testE2EUserUpdate(t *testing.T) {
	newToken := os.Getenv("PIXELA4GO_USER_SECOND_TOKEN")
	input := &UserUpdateInput{
		NewToken: String(newToken),
	}
	result, err := e2eClient.User().Update(input)
	if err != nil {
		t.Errorf("User.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("User.Update() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("User.Update() got: %+v\nwant: %d", result, http.StatusOK)
	}

	e2eClient.Token = newToken
}

func testE2EUserDelete(t *testing.T) {
	result, err := e2eClient.User().Delete()
	if err != nil {
		t.Errorf("User.Delete() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("User.Delete() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("User.Delete() got: %+v\nwant: %d", result, http.StatusOK)
	}
}
