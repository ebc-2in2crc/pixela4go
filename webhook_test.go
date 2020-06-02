package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWebhook_CreateCreateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &WebhookCreateInput{
		GraphID: String(graphID),
		Type:    String(WebhookTypeIncrement),
	}
	param, err := client.Webhook().createCreateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"graphID":"graph-id","type":"increment"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestWebhook_Create(t *testing.T) {
	s := `{"webhookHash":"webhook-hash","message":"Success.","isSuccess":true}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &WebhookCreateInput{
		GraphID: String(graphID),
		Type:    String(WebhookTypeIncrement),
	}
	result, err := client.Webhook().Create(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &WebhookCreateResult{
		WebhookHash: "webhook-hash",
		Result:      Result{Message: "Success.", IsSuccess: true},
	}
	if *result != *expect {
		t.Errorf("got: %v\nwant: %v", result, expect)
	}
}

func TestWebhook_CreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &WebhookCreateInput{
		GraphID: String(graphID),
		Type:    String(WebhookTypeIncrement),
	}
	result, err := client.Webhook().Create(input)

	testAPIFailedResult(t, &result.Result, err)
}

func TestWebhook_CreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &WebhookCreateInput{
		GraphID: String(graphID),
		Type:    String(WebhookTypeIncrement),
	}
	_, err := client.Webhook().Create(input)

	testPageNotFoundError(t, err)
}

func TestCreateWebhook_GetAllRequestParameter(t *testing.T) {
	client := New(userName, token)
	param, err := client.Webhook().createGetAllRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks", userName)
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

func TestWebhook_GetAll(t *testing.T) {
	s := `{"webhooks":[{"webhookHash":"webhook-hash","graphID":"test-graph","type":"increment"}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	definitions, err := client.Webhook().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &WebhookDefinitions{
		Webhooks: []WebhookDefinition{
			{
				WebhookHash: "webhook-hash",
				GraphID:     "test-graph",
				Type:        "increment",
			},
		},
		Result: Result{IsSuccess: true},
	}
	if reflect.DeepEqual(definitions, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestWebhook_GetAllFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	result, err := client.Webhook().GetAll()

	testAPIFailedResult(t, &result.Result, err)
}

func TestWebhook_GetAllError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	_, err := client.Webhook().GetAll()

	testPageNotFoundError(t, err)
}

func TestWebhook_CreateDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &WebhookDeleteInput{WebhookHash: String("webhook-hash")}
	param, err := client.Webhook().createDeleteRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks/webhook-hash", userName)
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

func TestWebhook_Delete(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &WebhookDeleteInput{WebhookHash: String("webhook-hash")}
	result, err := client.Webhook().Delete(input)

	testSuccess(t, result, err)
}

func TestWebhook_DeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &WebhookDeleteInput{WebhookHash: String("webhook-hash")}
	result, err := client.Webhook().Delete(input)

	testAPIFailedResult(t, result, err)
}

func TestWebhook_DeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &WebhookDeleteInput{WebhookHash: String("webhook-hash")}
	_, err := client.Webhook().Delete(input)

	testPageNotFoundError(t, err)
}

func TestWebhook_CreateInvokeRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &WebhookInvokeInput{WebhookHash: String("webhook-hash")}
	param, err := client.Webhook().createInvokeRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURL+"/users/%s/webhooks/webhook-hash", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[contentLength] != "0" {
		t.Errorf("%s: %s\nwant: %s", contentLength, param.Header[contentLength], "0")
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestWebhook_Invoke(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &WebhookInvokeInput{WebhookHash: String("webhook-hash")}
	result, err := client.Webhook().Invoke(input)

	testSuccess(t, result, err)
}

func TestWebhook_InvokeFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &WebhookInvokeInput{WebhookHash: String("webhook-hash")}
	result, err := client.Webhook().Invoke(input)

	testAPIFailedResult(t, result, err)
}

func TestWebhook_InvokeError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &WebhookInvokeInput{WebhookHash: String("webhook-hash")}
	_, err := client.Webhook().Invoke(input)

	testPageNotFoundError(t, err)
}
