package pixela

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// APIBaseURL is Base URL for API requests.
const APIBaseURL = "https://pixe.la"

// APIBaseURLForV1 is Base URL for API version 1 requests.
const APIBaseURLForV1 = APIBaseURL + "/v1"
const (
	contentType   = "Content-Type"
	contentLength = "Content-Length"
	userToken     = "X-USER-TOKEN"
)

type requestParameter struct {
	Method string
	URL    string
	Header map[string]string
	Body   []byte
}

// Result is Pixela API Result struct.
type Result struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}

func newHTTPRequest(param *requestParameter) (*http.Request, error) {
	req, err := http.NewRequest(
		param.Method,
		param.URL,
		bytes.NewReader(param.Body))
	if err != nil {
		return &http.Request{}, errors.Wrap(err, "failed to create http.Request")
	}
	if param.Header != nil {
		for k, v := range param.Header {
			req.Header.Add(k, v)
		}
	}
	req.Header.Set(contentType, "application/json")

	return req, nil
}

func doRequest(param *requestParameter) ([]byte, error) {
	req, err := newHTTPRequest(param)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to create http.Request")
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if clientMock != nil {
		resp, err = clientMock.do(req)
	}
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed http.Client do")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to read response body")
	}

	return b, nil
}

func mustDoRequest(param *requestParameter) ([]byte, error) {
	req, err := newHTTPRequest(param)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed to create http.Request")
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if clientMock != nil {
		resp, err = clientMock.do(req)
	}
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed http.Client do")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to read response body")
	}

	if resp.StatusCode >= 300 {
		return b, errors.Errorf("failed to call API: %s", string(b))
	}

	return b, nil
}

func doRequestAndParseResponse(param *requestParameter) (*Result, error) {
	req, err := newHTTPRequest(param)
	if err != nil {
		return &Result{}, errors.Wrap(err, "failed to create http.Request")
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if clientMock != nil {
		resp, err = clientMock.do(req)
	}
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed http.Client do")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to read response body")
	}

	return parseNormalResponse(b)
}

func parseNormalResponse(b []byte) (*Result, error) {
	var result Result
	if err := json.Unmarshal(b, &result); err != nil {
		return &Result{}, errors.Wrapf(err, "failed to unmarshal json: "+string(b))
	}
	return &result, nil
}
