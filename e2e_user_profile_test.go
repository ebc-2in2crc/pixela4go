package pixela

import (
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
}
