package pixela

import (
	"net/http"
	"testing"
)

func testE2EUserProfileUpdate(t *testing.T) {
	input := &UserProfileUpdateInput{
		DisplayName: String("displayName"),
	}
	result, err := e2eClient.UserProfile().Update(input)
	if err != nil {
		t.Errorf("UserProfile.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("UserProfile.Update() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("UserProfile.UPdate() got: %+v\nwant: %d", result, http.StatusOK)
	}
}
