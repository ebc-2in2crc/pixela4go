package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var detail = &SlackDetail{
	URL:         String("slack-url"),
	UserName:    String("slack-user"),
	ChannelName: String("slack-channel-name"),
}

func TestChannel_CreateCreateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &ChannelCreateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	param, err := client.Channel().createCreateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"id":"channel-id","name":"channel-name","type":"slack","detail":{"url":"slack-url","userName":"slack-user","channelName":"slack-channel-name"}}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestChannel_Create(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &ChannelCreateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	result, err := client.Channel().Create(input)

	testSuccess(t, result, err)
}

func TestChannel_CreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &ChannelCreateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	result, err := client.Channel().Create(input)

	testAPIFailedResult(t, result, err)
}

func TestChannel_CreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &ChannelCreateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	_, err := client.Channel().Create(input)

	testPageNotFoundError(t, err)
}

func TestChannel_CreateGetRequestParameter(t *testing.T) {
	client := New(userName, token)
	param, err := client.Channel().createGetRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels", userName)
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

func TestChannel_GetAll(t *testing.T) {
	s := `{"channels":[{"id":"channel-id","name":"channel-name","type":"slack","detail":{"url":"slack-url","userName":"slack-user","channelName":"slack-channel-name"}}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	definitions, err := client.Channel().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if len(definitions.Channels) != 1 {
		t.Errorf("got: %d definitions \nwant: one", len(definitions.Channels))
	}

	channel := definitions.Channels[0]
	if channel.ID != "channel-id" {
		t.Errorf("got: %v\nwant: %v", channel.ID, "channel-id")
	}
	if channel.Name != "channel-name" {
		t.Errorf("got: %v\nwant: %v", channel.Name, "channel-name")
	}
	if channel.Type != ChannelTypeSlack {
		t.Errorf("got: %v\nwant: %v", channel.Type, ChannelTypeSlack)
	}

	d := channel.Detail.(SlackDetail)
	expect := SlackDetail{
		URL:         String("slack-url"),
		UserName:    String("slack-user"),
		ChannelName: String("slack-channel-name"),
	}
	if reflect.DeepEqual(d, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestChannel_GetAllFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	result, err := client.Channel().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestChannel_GetAllError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	_, err := client.Channel().GetAll()

	testPageNotFoundError(t, err)
}

func TestChannel_CreateUpdateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &ChannelUpdateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	param, err := client.Channel().createUpdateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels/channel-id", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"name":"channel-name","type":"slack","detail":{"url":"slack-url","userName":"slack-user","channelName":"slack-channel-name"}}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestChannel_Update(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &ChannelUpdateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	result, err := client.Channel().Update(input)

	testSuccess(t, result, err)
}

func TestChannel_UpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &ChannelUpdateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	result, err := client.Channel().Update(input)

	testAPIFailedResult(t, result, err)
}

func TestChannel_UpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &ChannelUpdateInput{
		ID:          String("channel-id"),
		Name:        String("channel-name"),
		Type:        String(ChannelTypeSlack),
		SlackDetail: detail,
	}
	_, err := client.Channel().Update(input)

	testPageNotFoundError(t, err)
}

func TestChannel_CreateDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &ChannelDeleteInput{ID: String("channel-id")}
	param, err := client.Channel().createDeleteRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/channels/channel-id", userName)
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

func TestChannel_Delete(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &ChannelDeleteInput{ID: String("channel-id")}
	result, err := client.Channel().Delete(input)

	testSuccess(t, result, err)
}

func TestChannel_DeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &ChannelDeleteInput{ID: String("channel-id")}
	result, err := client.Channel().Delete(input)

	testAPIFailedResult(t, result, err)
}

func TestChannel_DeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &ChannelDeleteInput{ID: String("channel-id")}
	_, err := client.Channel().Delete(input)

	testPageNotFoundError(t, err)
}
