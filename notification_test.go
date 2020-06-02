package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNotification_CreateCreateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &NotificationCreateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}

	param, err := client.Notification().createCreateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"id":"notification-id","name":"notification-name","target":"quantity","condition":"\u003e","threshold":"3","remindBy":"23","channelID":"channel-id"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestNotification_Create(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &NotificationCreateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	result, err := client.Notification().Create(input)

	testSuccess(t, result, err)
}

func TestNotification_CreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &NotificationCreateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	result, err := client.Notification().Create(input)

	testAPIFailedResult(t, result, err)
}

func TestNotification_CreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &NotificationCreateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	_, err := client.Notification().Create(input)

	testPageNotFoundError(t, err)
}

func TestNotification_CreateGetRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &NotificationGetAllInput{GraphID: String(graphID)}
	param, err := client.Notification().createGetRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", userName, graphID)
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

func TestNotification_GetAll(t *testing.T) {
	s := `{"notifications":[{"id":"notification-id","name":"notification-name","target":"quantity","condition":"\u003e","threshold":"3","remindBy":"23","channelID":"channel-id"}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &NotificationGetAllInput{GraphID: String(graphID)}
	definitions, err := client.Notification().GetAll(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &NotificationDefinitions{
		Notifications: []NotificationDefinition{
			{
				ID:        "notification-id",
				Name:      "notification-name",
				Target:    NotificationTargetQuantity,
				Condition: NotificationConditionGreaterThan,
				Threshold: "3",
				RemindBy:  "23",
				ChannelID: "channel-id",
			},
		},
		Result: Result{IsSuccess: true},
	}

	if reflect.DeepEqual(definitions, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestNotification_CreateUpdateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &NotificationUpdateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	param, err := client.Notification().createUpdateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", userName, graphID, "notification-id")
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"name":"notification-name","target":"quantity","condition":"\u003e","threshold":"3","remindBy":"23","channelID":"channel-id"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestNotification_Update(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &NotificationUpdateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	result, err := client.Notification().Update(input)

	testSuccess(t, result, err)
}

func TestNotification_UpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &NotificationUpdateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	result, err := client.Notification().Update(input)

	testAPIFailedResult(t, result, err)
}

func TestNotification_UpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &NotificationUpdateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	_, err := client.Notification().Update(input)

	testPageNotFoundError(t, err)
}

func TestNotification_CreateDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &NotificationDeleteInput{GraphID: String(graphID), ID: String("notification-id")}
	param, err := client.Notification().createDeleteRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", userName, graphID, "notification-id")
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

func TestNotification_Delete(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &NotificationDeleteInput{GraphID: String(graphID), ID: String("notification-id")}
	result, err := client.Notification().Delete(input)

	testSuccess(t, result, err)
}

func TestNotification_DeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &NotificationDeleteInput{GraphID: String(graphID), ID: String("notification-id")}
	result, err := client.Notification().Delete(input)

	testAPIFailedResult(t, result, err)
}

func TestNotification_DeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &NotificationDeleteInput{GraphID: String(graphID), ID: String("notification-id")}
	_, err := client.Notification().Delete(input)

	testPageNotFoundError(t, err)
}
