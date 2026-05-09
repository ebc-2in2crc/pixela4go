package pixela

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		return &http.Request{}, fmt.Errorf("failed to create http.Request: %w", err)
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
			m.err = fmt.Errorf("failed to create http.Request: %w", err)
			return
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if clientMock != nil {
			resp, err = clientMock.do(req)
		}
		if err != nil {
			m.err = fmt.Errorf("failed http.Client do: %w", err)
			return
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			m.err = fmt.Errorf("failed to read response body: %w", err)
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
		return []byte{}, fmt.Errorf("failed to create http.Request: %w", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if clientMock != nil {
		resp, err = clientMock.do(req)
	}
	if err != nil {
		return []byte{}, fmt.Errorf("failed http.Client do: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 300 {
		return b, fmt.Errorf("failed to call API: %s", string(b))
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
		return &Result{}, fmt.Errorf("failed to parse normal response: %w", err)
	}

	r.StatusCode = retry.statusCode
	return r, nil
}

func parseNormalResponse(b []byte) (*Result, error) {
	var result Result
	if err := json.Unmarshal(b, &result); err != nil {
		return &Result{}, fmt.Errorf("failed to unmarshal json: %s: %w", string(b), err)
	}
	return &result, nil
}
