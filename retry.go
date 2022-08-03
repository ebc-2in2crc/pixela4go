package pixela

import (
	"math"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const maxRetryCount = 20

// RetryCount number of retries when an API call is rejected (max: 20)
var RetryCount = 0

// ErrAPICallRejected API call rejected.
// See: https://help.pixe.la/en/blog/release-request-rejecting
var ErrAPICallRejected = errors.New("api call rejected")

type retryer struct {
	processFunc func(r *retryer)
	maxRetry    int
	statusCode  int
	body        []byte
	err         error
}

func (m *retryer) do() error {
	for i := 0; i <= m.maxRetry; i++ {
		m.process()
		if !m.shouldRetry() {
			return m.err
		}

		waitTime := m.getWaitTimeExp(i, 100)
		time.Sleep(time.Millisecond * time.Duration(waitTime))
	}

	return ErrAPICallRejected
}

func (m *retryer) process() {
	m.processFunc(m)
}

func (m *retryer) shouldRetry() bool {
	if m.err != nil {
		return false
	}
	if m.statusCode != http.StatusServiceUnavailable {
		return false
	}

	r, err := parseNormalResponse(m.body)
	if err != nil {
		m.err = errors.Wrapf(err, "failed to parse normal response")
		return false
	}

	return r.IsRejected
}

func (m *retryer) getWaitTimeExp(retryCount int, baseDelayMilliSeconds int) (delayMilliSeconds int) {
	if retryCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(retryCount)) * float64(baseDelayMilliSeconds))
}

func getRetryCount() int {
	if RetryCount < 0 {
		return 0
	}
	if RetryCount > maxRetryCount {
		return maxRetryCount
	}
	return RetryCount
}
