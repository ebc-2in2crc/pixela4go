package pixela

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	userName = "user"
	token    = "token"
	graphID  = "graph-id"
)

type httpClientMock struct {
	statusCode int
	body       []byte
}

func (c *httpClientMock) do(req *http.Request) (*http.Response, error) {
	resp := &http.Response{}
	resp.StatusCode = c.statusCode
	resp.Body = ioutil.NopCloser(bytes.NewReader(c.body))
	return resp, nil
}

var clientMock *httpClientMock

func newOKMock() *httpClientMock {
	return &httpClientMock{
		statusCode: http.StatusOK,
		body:       []byte(`{"message":"Success.","isSuccess":true}`),
	}
}

func testSuccess(t *testing.T, actual *Result, err error) {
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Result{
		Message:   "Success.",
		IsSuccess: true,
	}
	if *actual != *expect {
		t.Errorf("got: %v\nwant: %v", actual, expect)
	}
}

func newAPIFailedMock() *httpClientMock {
	return &httpClientMock{
		statusCode: http.StatusNotFound,
		body:       []byte(`{"message":"failed.","isSuccess":false}`),
	}
}

func testAPIFailedResult(t *testing.T, result *Result, err error) {
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Result{
		Message:   "failed.",
		IsSuccess: false,
	}
	if *result != *expect {
		t.Errorf("got: %v\nwant: %v", result, expect)
	}
}

func newPageNotFoundMock() *httpClientMock {
	return &httpClientMock{
		statusCode: http.StatusNotFound,
		body:       []byte("404 page not found"),
	}
}

func testPageNotFoundError(t *testing.T, err error) {
	expect := "failed to unmarshal json:"
	if err == nil {
		t.Errorf("got: nil\nwant: %s", expect)
	}

	if err != nil && strings.HasPrefix(err.Error(), expect) == false {
		t.Errorf("got: %s\nwant: %s", err.Error(), expect)
	}
}
