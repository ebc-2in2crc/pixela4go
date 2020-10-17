package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestPixel_CreateCreateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelCreateInput{
		Date:         String("20180915"),
		Quantity:     String("5"),
		OptionalData: String("{\"key\":\"value\"}"),
		GraphID:      String(graphID),
	}
	param, err := client.Pixel().createCreateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"date":"20180915","quantity":"5","optionalData":"{\"key\":\"value\"}"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestPixel_Create(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &PixelCreateInput{Date: String("20180915"), Quantity: String("5"), GraphID: String(graphID)}
	result, err := client.Pixel().Create(input)

	testSuccess(t, result, err)
}

func TestPixel_CreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &PixelCreateInput{Date: String("20180915"), Quantity: String("5"), GraphID: String(graphID)}
	result, err := client.Pixel().Create(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_CreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &PixelCreateInput{Date: String("20180915"), Quantity: String("5"), GraphID: String(graphID)}
	_, err := client.Pixel().Create(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateIncrementRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelIncrementInput{GraphID: String(graphID)}
	param, err := client.Pixel().createIncrementRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/increment", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[contentLength] != "0" {
		t.Errorf("%s: %s\nwant: %s", contentLength, param.Header[contentLength], "0")
	}
	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestPixel_Increment(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &PixelIncrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Increment(input)

	testSuccess(t, result, err)
}

func TestPixel_IncrementFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &PixelIncrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Increment(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_IncrementError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &PixelIncrementInput{GraphID: String(graphID)}
	_, err := client.Pixel().Increment(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateDecrementRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelDecrementInput{GraphID: String(graphID)}
	param, err := client.Pixel().createDecrementRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/decrement", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[contentLength] != "0" {
		t.Errorf("%s: %s\nwant: %s", contentLength, param.Header[contentLength], "0")
	}
	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}
}

func TestPixel_Decrement(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &PixelDecrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Decrement(input)

	testSuccess(t, result, err)
}

func TestPixel_DecrementFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &PixelDecrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Decrement(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_DecrementError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &PixelDecrementInput{GraphID: String(graphID)}
	_, err := client.Pixel().Decrement(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateGetRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	param, err := client.Pixel().createGetRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/%s", userName, graphID, "20180915")
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

func TestPixel_Get(t *testing.T) {
	s := `{"quantity": "5","optionalData":"{\"key\":\"value\"}"}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	quantity, err := client.Pixel().Get(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Quantity{
		Quantity:     "5",
		OptionalData: "{\"key\":\"value\"}",
		Result:       Result{IsSuccess: true},
	}
	if *quantity != *expect {
		t.Errorf("got: %v\nwant: %v", quantity, expect)
	}
}

func TestPixel_GetFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	result, err := client.Pixel().Get(input)

	testAPIFailedResult(t, &result.Result, err)
}

func TestPixel_GetError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	_, err := client.Pixel().Get(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateUpdateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelUpdateInput{
		GraphID:      String(graphID),
		Date:         String("20180915"),
		Quantity:     String("5"),
		OptionalData: String("{\"key\":\"value\"}"),
	}
	param, err := client.Pixel().createUpdateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/20180915", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"quantity":"5","optionalData":"{\"key\":\"value\"}"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestPixel_Update(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &PixelUpdateInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	result, err := client.Pixel().Update(input)

	testSuccess(t, result, err)
}

func TestPixel_UpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &PixelUpdateInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	result, err := client.Pixel().Update(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_UpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &PixelUpdateInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	_, err := client.Pixel().Update(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	param, err := client.Pixel().createDeleteRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/20180915", userName, graphID)
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

func TestPixel_Delete(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	result, err := client.Pixel().Delete(input)

	testSuccess(t, result, err)
}

func TestPixel_DeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	result, err := client.Pixel().Delete(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_DeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	_, err := client.Pixel().Delete(input)

	testPageNotFoundError(t, err)
}
