package pixela

import (
	"bytes"
	"context"
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
	Message    string `json:"message"`
	IsSuccess  bool   `json:"isSuccess"`
	IsRejected bool   `json:"isRejected"`
	StatusCode int    `json:"statusCode"`
}

func newHTTPRequest(ctx context.Context, param *requestParameter) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
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

func doRequest(ctx context.Context, param *requestParameter) ([]byte, int, error) {
	retry := retryer{
		processFunc: processFunc(ctx, param),
		maxRetry:    getRetryCount(),
	}
	if err := retry.do(ctx); err != nil {
		return []byte{}, 0, err
	}

	return retry.body, retry.statusCode, nil
}

func processFunc(ctx context.Context, param *requestParameter) func(m *retryer) {
	return func(m *retryer) {
		req, err := newHTTPRequest(ctx, param)
		if err != nil {
			m.err = errors.Wrap(err, "failed to create http.Request")
			return
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if clientMock != nil {
			resp, err = clientMock.do(req)
		}
		if err != nil {
			m.err = errors.Wrapf(err, "failed http.Client do")
			return
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			m.err = errors.Wrapf(err, "failed to read response body")
			return
		}

		m.body = b
		m.statusCode = resp.StatusCode
		m.err = nil
	}
}

func mustDoRequest(ctx context.Context, param *requestParameter) ([]byte, error) {
	req, err := newHTTPRequest(ctx, param)
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

func doRequestAndParseResponse(ctx context.Context, param *requestParameter) (*Result, error) {
	retry := retryer{
		processFunc: processFunc(ctx, param),
		maxRetry:    getRetryCount(),
	}
	if err := retry.do(ctx); err != nil {
		return &Result{}, err
	}

	r, err := parseNormalResponse(retry.body)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to parse normal response")
	}

	r.StatusCode = retry.statusCode
	return r, nil
}

func parseNormalResponse(b []byte) (*Result, error) {
	var result Result
	if err := json.Unmarshal(b, &result); err != nil {
		return &Result{}, errors.Wrapf(err, "failed to unmarshal json: "+string(b))
	}
	return &result, nil
}
