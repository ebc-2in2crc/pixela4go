package pixela

import (
	"fmt"
	"os"
	"testing"
)

var e2eClient *Client

func TestE2E(t *testing.T) {
	if os.Getenv("PIXELA4GO_E2E_TEST_RUN") == "" {
		msg := `E2E test skip.
If you run E2E test, Set below environment variables.

- PIXELA4GO_E2E_TEST_RUN=ON
- PIXELA4GO_USER_NAME=<pixela-username-for-testing>
- PIXELA4GO_USER_FIRST_TOKEN=<pixela-user-token-for-testing>
- PIXELA4GO_USER_SECOND_TOKEN=<pixela-user-token-for-testing>`
		fmt.Println(msg)
		return
	}

	initE2ETest()

	testE2EUserCreate(t)
	testE2EUserUpdate(t)

	testE2EUserProfileUpdate(t)

	testE2EGraphCreate(t)
	testE2EGraphUpdate(t)
	testE2EGraphGetAll(t)
	testE2EGraphStopwatch(t)
	testE2EGraphGetPixelDates(t)
	testE2EGraphGetSVG(t)
	testE2EGraphStats(t)

	testE2EPixelCreate(t)
	testE2EPixelIncrement(t)
	testE2EPixelDecrement(t)
	testE2EPixelUpdate(t)
	testE2EPixelGet(t)
	testE2EPixelDelete(t)

	testE2EWebhookCreate(t)
	testE2EWebhookInvoke(t)
	testE2EWebhookGetAll(t)
	testE2EWebhookDelete(t)

	testE2EGraphDelete(t)
	testE2EUserDelete(t)
}

func initE2ETest() {
	clientMock = nil

	name := os.Getenv("PIXELA4GO_USER_NAME")
	token := os.Getenv("PIXELA4GO_USER_FIRST_TOKEN")
	e2eClient = New(name, token)
}
