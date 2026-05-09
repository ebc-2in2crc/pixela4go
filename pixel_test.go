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
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelCreateInput{Date: String("20180915"), Quantity: String("5"), GraphID: String(graphID)}
	result, err := client.Pixel().Create(input)

	testSuccess(t, result, err)
}

func TestPixel_CreateFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelCreateInput{Date: String("20180915"), Quantity: String("5"), GraphID: String(graphID)}
	result, err := client.Pixel().Create(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_CreateError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelCreateInput{Date: String("20180915"), Quantity: String("5"), GraphID: String(graphID)}
	_, err := client.Pixel().Create(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateIncrementRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelIncrementInput{GraphID: String(graphID)}
	param := client.Pixel().createIncrementRequestParameter(input)

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
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelIncrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Increment(input)

	testSuccess(t, result, err)
}

func TestPixel_IncrementFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelIncrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Increment(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_IncrementError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelIncrementInput{GraphID: String(graphID)}
	_, err := client.Pixel().Increment(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateDecrementRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelDecrementInput{GraphID: String(graphID)}
	param := client.Pixel().createDecrementRequestParameter(input)

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
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelDecrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Decrement(input)

	testSuccess(t, result, err)
}

func TestPixel_DecrementFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelDecrementInput{GraphID: String(graphID)}
	result, err := client.Pixel().Decrement(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_DecrementError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelDecrementInput{GraphID: String(graphID)}
	_, err := client.Pixel().Decrement(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateGetRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	param := client.Pixel().createGetRequestParameter(input)

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
	client := New(userName, token)
	client.HTTPClient = &httpClientMock{statusCode: http.StatusOK, body: b}
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	quantity, err := client.Pixel().Get(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Quantity{
		Quantity:     "5",
		OptionalData: "{\"key\":\"value\"}",
		Result:       Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if *quantity != *expect {
		t.Errorf("got: %v\nwant: %v", quantity, expect)
	}
}

func TestPixel_GetFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelGetInput{GraphID: String(graphID), Date: String("20180915")}
	result, err := client.Pixel().Get(input)

	testAPIFailedResult(t, &result.Result, err)
}

func TestPixel_GetError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
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
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelUpdateInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	result, err := client.Pixel().Update(input)

	testSuccess(t, result, err)
}

func TestPixel_UpdateFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelUpdateInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	result, err := client.Pixel().Update(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_UpdateError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelUpdateInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	_, err := client.Pixel().Update(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateAddRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelAddInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("5"),
	}
	param, err := client.Pixel().createAddRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/20180915/add", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"quantity":"5"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestPixel_Add(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelAddInput{GraphID: String(graphID), Date: String("20180915"), Quantity: String("5")}
	result, err := client.Pixel().Add(input)

	testSuccess(t, result, err)
}

func TestPixel_AddFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelAddInput{GraphID: String(graphID), Date: String("20180915"), Quantity: String("5")}
	result, err := client.Pixel().Add(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_AddError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelAddInput{GraphID: String(graphID), Date: String("20180915"), Quantity: String("5")}
	_, err := client.Pixel().Add(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateSubtractRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelSubtractInput{
		GraphID:  String(graphID),
		Date:     String("20180915"),
		Quantity: String("3"),
	}
	param, err := client.Pixel().createSubtractRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/20180915/subtract", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"quantity":"3"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestPixel_Subtract(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelSubtractInput{GraphID: String(graphID), Date: String("20180915"), Quantity: String("3")}
	result, err := client.Pixel().Subtract(input)

	testSuccess(t, result, err)
}

func TestPixel_SubtractFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelSubtractInput{GraphID: String(graphID), Date: String("20180915"), Quantity: String("3")}
	result, err := client.Pixel().Subtract(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_SubtractError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelSubtractInput{GraphID: String(graphID), Date: String("20180915"), Quantity: String("3")}
	_, err := client.Pixel().Subtract(input)

	testPageNotFoundError(t, err)
}

func TestPixel_CreateDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	param := client.Pixel().createDeleteRequestParameter(input)

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
	client := New(userName, token)
	client.HTTPClient = newOKMock()
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	result, err := client.Pixel().Delete(input)

	testSuccess(t, result, err)
}

func TestPixel_DeleteFail(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newAPIFailedMock()
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	result, err := client.Pixel().Delete(input)

	testAPIFailedResult(t, result, err)
}

func TestPixel_DeleteError(t *testing.T) {
	client := New(userName, token)
	client.HTTPClient = newPageNotFoundMock()
	input := &PixelDeleteInput{GraphID: String(graphID), Date: String("20180915")}
	_, err := client.Pixel().Delete(input)

	testPageNotFoundError(t, err)
}
