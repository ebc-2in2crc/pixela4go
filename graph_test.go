package pixela

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestGraph_CreateCreateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphCreateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Type:                String(GraphTypeInt),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
	}
	param, err := client.Graph().createCreateRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs", userName)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"id":"graph-id","name":"name","unit":"times","type":"int","color":"shibafu","timezone":"UTC","selfSufficient":"increment","isSecret":true,"publishOptionalData":true}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraph_Create(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphCreateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Type:                String(GraphTypeInt),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
		StartOnMonday:       Bool(true),
	}
	result, err := client.Graph().Create(input)

	testSuccess(t, result, err)
}

func TestGraph_CreateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphCreateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Type:                String(GraphTypeInt),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
	}
	result, err := client.Graph().Create(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_CreateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphCreateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Type:                String(GraphTypeInt),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
	}
	_, err := client.Graph().Create(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateGetAllRequestParameter(t *testing.T) {
	client := New(userName, token)
	param, err := client.Graph().createGetAllRequestParameter()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs", userName)
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

func TestGraph_GetAll(t *testing.T) {
	s := `{"graphs":[{"id":"test-graph","name":"graph-name","unit":"commit","type":"int","color":"shibafu","timezone":"Asia/Tokyo","purgeCacheURLs":["https://camo.githubusercontent.com/xxx/xxxx"],"selfSufficient":"increment","isSecret":true,"publishOptionalData":true}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	definitions, err := client.Graph().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &GraphDefinitions{
		Graphs: []GraphDefinition{
			{
				ID:                  "test-graph",
				Name:                "graph-name",
				Unit:                "commit",
				Type:                "int",
				Color:               "shibafu",
				TimeZone:            "Asia/Tokyo",
				PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
				SelfSufficient:      "increment",
				IsSecret:            true,
				PublishOptionalData: true,
			},
		},
		Result: Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if reflect.DeepEqual(definitions, expect) == false {
		t.Errorf("got: %v\nwant: %v", definitions, expect)
	}
}

func TestGraph_GetAllFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	result, err := client.Graph().GetAll()
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraph_GetAllError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	_, err := client.Graph().GetAll()

	testPageNotFoundError(t, err)
}

func TestGraph_CreateGetLatestPixelRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphGetLatestPixelInput{ID: String(graphID)}
	param, err := client.Graph().createGetLatestPixelRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/latest", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := ""
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraph_GetLatestPixel(t *testing.T) {
	s := `{"date":"20240414","quantity":"5","optionalData":"{\"key\":\"value\"}"}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphGetLatestPixelInput{
		ID: String(graphID),
	}
	pixel, err := client.Graph().GetLatestPixel(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &GraphPixel{
		Date:         "20240414",
		Quantity:     "5",
		OptionalData: "{\"key\":\"value\"}",
		Result:       Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if reflect.DeepEqual(pixel, expect) == false {
		t.Errorf("got: %v\nwant: %v", pixel, expect)
	}
}

func TestGraph_GetLatestPixelError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphGetLatestPixelInput{
		ID: String(graphID),
	}
	_, err := client.Graph().GetLatestPixel(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateGetTodayRequestParameter(t *testing.T) {
	client := New(userName, token)

	// Test with ReturnEmpty = nil
	input1 := &GraphGetTodayInput{
		ID: String(graphID),
	}
	param1, err := client.Graph().createGetTodayRequestParameter(input1)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param1.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param1.Method, http.MethodGet)
	}

	expectURL1 := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/today", userName, graphID)
	if param1.URL != expectURL1 {
		t.Errorf("URL: %s\nwant: %s", param1.URL, expectURL1)
	}

	if param1.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param1.Header[userToken], token)
	}

	if bytes.Equal(param1.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param1.Body))
	}

	// Test with ReturnEmpty = true
	input2 := &GraphGetTodayInput{
		ID:          String(graphID),
		ReturnEmpty: Bool(true),
	}
	param2, err := client.Graph().createGetTodayRequestParameter(input2)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	// Parse the actual URL to compare parts independently of query parameter order
	actualURL, err := url.Parse(param2.URL)
	if err != nil {
		t.Errorf("Failed to parse actual URL: %v", err)
	}

	// Create the expected base URL
	expectedBaseURL := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/today", userName, graphID)

	// Check that the base URL matches
	actualBaseURL := actualURL.Scheme + "://" + actualURL.Host + actualURL.Path
	if actualBaseURL != expectedBaseURL {
		t.Errorf("Base URL: %s\nwant: %s", actualBaseURL, expectedBaseURL)
	}

	// Check that all expected query parameters are present with correct values
	query := actualURL.Query()
	if query.Get("returnEmpty") != "true" {
		t.Errorf("returnEmpty parameter: %s\nwant: %s", query.Get("returnEmpty"), "true")
	}

	// Check that there are no unexpected query parameters
	if len(query) != 1 {
		t.Errorf("Number of query parameters: %d\nwant: %d", len(query), 1)
	}
}

func TestGraph_GetToday(t *testing.T) {
	s := `{"date":"20240414","quantity":"5","optionalData":"{\"key\":\"value\"}"}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphGetTodayInput{
		ID: String(graphID),
	}
	pixel, err := client.Graph().GetToday(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &GraphPixel{
		Date:         "20240414",
		Quantity:     "5",
		OptionalData: "{\"key\":\"value\"}",
		Result:       Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if reflect.DeepEqual(pixel, expect) == false {
		t.Errorf("got: %v\nwant: %v", pixel, expect)
	}
}

func TestGraph_GetTodayError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphGetTodayInput{
		ID: String(graphID),
	}
	_, err := client.Graph().GetToday(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateGetSVGRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphGetSVGInput{
		ID:          String(graphID),
		Date:        String("20180101"),
		Mode:        String(GraphModeShort),
		Appearance:  String(GraphAppearanceDark),
		LessThan:    String("10"),
		GreaterThan: String("5"),
	}
	param, err := client.Graph().createGetSVGRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	// Parse the actual URL to compare parts independently of query parameter order
	actualURL, err := url.Parse(param.URL)
	if err != nil {
		t.Errorf("Failed to parse actual URL: %v", err)
	}

	// Create the expected base URL
	expectedBaseURL := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", userName, graphID)

	// Check that the base URL matches
	actualBaseURL := actualURL.Scheme + "://" + actualURL.Host + actualURL.Path
	if actualBaseURL != expectedBaseURL {
		t.Errorf("Base URL: %s\nwant: %s", actualBaseURL, expectedBaseURL)
	}

	// Check that all expected query parameters are present with correct values
	query := actualURL.Query()
	if query.Get("date") != "20180101" {
		t.Errorf("date parameter: %s\nwant: %s", query.Get("date"), "20180101")
	}
	if query.Get("mode") != "short" {
		t.Errorf("mode parameter: %s\nwant: %s", query.Get("mode"), "short")
	}
	if query.Get("appearance") != "dark" {
		t.Errorf("appearance parameter: %s\nwant: %s", query.Get("appearance"), "dark")
	}
	if query.Get("lessThan") != "10" {
		t.Errorf("lessThan parameter: %s\nwant: %s", query.Get("lessThan"), "10")
	}
	if query.Get("greaterThan") != "5" {
		t.Errorf("greaterThan parameter: %s\nwant: %s", query.Get("greaterThan"), "5")
	}

	// Check that there are no unexpected query parameters
	if len(query) != 5 {
		t.Errorf("Number of query parameters: %d\nwant: %d", len(query), 5)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestGraph_GetSVG(t *testing.T) {
	s := `<svg></svg>`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphGetSVGInput{
		ID:   String(graphID),
		Date: String("20180101"),
		Mode: String(GraphModeShort),
	}
	svg, err := client.Graph().GetSVG(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := s
	if svg != expect {
		t.Errorf("got: %s\nwant: %s", svg, expect)
	}
}

func TestGraph_GetSVGFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphGetSVGInput{
		ID:   String(graphID),
		Date: String("20180101"),
		Mode: String(GraphModeShort),
	}
	_, err := client.Graph().GetSVG(input)
	expect := "failed to do request: failed to call API: " + string(clientMock.body)
	if err == nil {
		t.Errorf("got: nil\nwant: %s", expect)
	}

	if err != nil && err.Error() != expect {
		t.Errorf("got: %s\nwant: %s", err.Error(), expect)
	}
}

func TestGraph_CreateUpdatePixelsRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphUpdatePixelsInput{
		ID: String(graphID),
		Pixels: []PixelInput{
			{
				Date:         String("20180101"),
				Quantity:     String("1"),
				OptionalData: nil,
			},
			{
				Date:         String("20180102"),
				Quantity:     String("2"),
				OptionalData: String("{\"key\":\"value\"}"),
			},
		},
	}
	param, err := client.Graph().createUpdatePixelsRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/pixels", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `[{"date":"20180101","quantity":"1"},{"date":"20180102","quantity":"2","optionalData":"{\"key\":\"value\"}"}]`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraph_UpdatePixels(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphUpdatePixelsInput{
		ID:     String(graphID),
		Pixels: []PixelInput{},
	}
	result, err := client.Graph().UpdatePixels(input)

	testSuccess(t, result, err)
}

func TestGraph_UpdatePixelsFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphUpdatePixelsInput{
		ID:     String(graphID),
		Pixels: []PixelInput{},
	}
	result, err := client.Graph().UpdatePixels(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_UpdatePixelsError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphUpdatePixelsInput{
		ID:     String(graphID),
		Pixels: []PixelInput{},
	}
	_, err := client.Graph().UpdatePixels(input)

	testPageNotFoundError(t, err)
}

func TestGraph_URL(t *testing.T) {
	client := New(userName, token)
	baseURL := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s.html", userName, graphID)
	params := []struct {
		mode   string
		expect string
	}{
		{mode: "", expect: baseURL},
		{mode: "simple", expect: baseURL + "?mode=simple"},
		{mode: "simple-short", expect: baseURL + "?mode=simple-short"},
		{mode: "badge", expect: baseURL + "?mode=badge"},
	}

	for _, p := range params {
		input := &GraphURLInput{ID: String(graphID), Mode: String(p.mode)}
		url := client.Graph().URL(input)
		if url != p.expect {
			t.Errorf("got: %s\nwant: %s", url, p.expect)
		}
	}
}

func TestGraph_CreateStatsRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphStatsInput{ID: String(graphID)}
	param, err := client.Graph().createStatsRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/stats", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if bytes.Equal(param.Body, []byte{}) == false {
		t.Errorf("Body: %s\nwant: \"\"", string(param.Body))
	}
}

func TestGraph_Stats(t *testing.T) {
	s := `{"totalPixelsCount":1,"maxQuantity":2,"maxDate":"2023-09-01","minQuantity":3,"minDate":"2023-09-02","totalQuantity":4,"avgQuantity":5.0,"todaysQuantity":6,"yesterdayQuantity":66}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphStatsInput{ID: String(graphID)}
	stats, err := client.Graph().Stats(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Stats{
		TotalPixelsCount:  1,
		MaxQuantity:       2,
		MaxDate:           "2023-09-01",
		MinQuantity:       3,
		MinDate:           "2023-09-02",
		TotalQuantity:     4,
		AvgQuantity:       5.0,
		TodaysQuantity:    6,
		YesterdayQuantity: 66,
		Result:            Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if *stats != *expect {
		t.Errorf("got: %v\nwant: %v", stats, expect)
	}
}

func TestGraph_StatsFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphStatsInput{ID: String(graphID)}
	result, err := client.Graph().Stats(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraph_StatsError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphStatsInput{ID: String(graphID)}
	_, err := client.Graph().Stats(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateUpdateRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphUpdateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
	}
	param, err := client.Graph().createUpdateRequestParameter(input)

	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"name":"name","unit":"times","color":"shibafu","timezone":"UTC","purgeCacheURLs":["https://camo.githubusercontent.com/xxx/xxxx"],"selfSufficient":"increment","isSecret":true,"publishOptionalData":true}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraph_Update(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphUpdateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
		StartOnMonday:       Bool(true),
	}
	result, err := client.Graph().Update(input)

	testSuccess(t, result, err)
}

func TestGraph_UpdateFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphUpdateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
	}
	result, err := client.Graph().Update(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_UpdateError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphUpdateInput{
		ID:                  String(graphID),
		Name:                String("name"),
		Unit:                String("times"),
		Color:               String(GraphColorShibafu),
		TimeZone:            String("UTC"),
		PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
		SelfSufficient:      String(GraphSelfSufficientIncrement),
		IsSecret:            Bool(true),
		PublishOptionalData: Bool(true),
	}
	_, err := client.Graph().Update(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateDeleteRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphDeleteInput{ID: String(graphID)}
	param, err := client.Graph().createDeleteRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodDelete {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodDelete)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", userName, graphID)
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

func TestGraph_Delete(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphDeleteInput{ID: String(graphID)}
	result, err := client.Graph().Delete(input)

	testSuccess(t, result, err)
}

func TestGraph_DeleteFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphDeleteInput{ID: String(graphID)}
	result, err := client.Graph().Delete(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_DeleteError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphDeleteInput{ID: String(graphID)}
	_, err := client.Graph().Delete(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateGetPixelDatesRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphGetPixelDatesInput{
		ID:       String(graphID),
		From:     String("20180101"),
		To:       String("20181231"),
		WithBody: Bool(true),
	}
	param, err := client.Graph().createGetPixelDatesRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/pixels?from=20180101&to=20181231&withBody=true", userName, graphID)
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

func TestGraph_GetPixelDates(t *testing.T) {
	s := `{"pixels":["20180101","20180331"]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphGetPixelDatesInput{ID: String(graphID), From: String("20180101"), To: String("20181231")}
	pixels, err := client.Graph().GetPixelDates(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Pixels{
		Pixels: []string{"20180101", "20180331"},
		Result: Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if reflect.DeepEqual(pixels, expect) == false {
		t.Errorf("got: %v\nwant: %v", pixels, expect)
	}
}

func TestGraph_GetPixelDatesWithBody(t *testing.T) {
	s := `{"pixels":[{"date":"20180331","quantity":"1","optionalData":"{\"key\":\"value\"}"}]}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphGetPixelDatesInput{
		ID:       String(graphID),
		From:     String("20180101"),
		To:       String("20181231"),
		WithBody: Bool(true),
	}
	pixels, err := client.Graph().GetPixelDates(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &Pixels{
		Pixels: []PixelWithBody{
			{
				Date:         "20180331",
				Quantity:     "1",
				OptionalData: "{\"key\":\"value\"}",
			},
		},
		Result: Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if reflect.DeepEqual(pixels, expect) == false {
		t.Errorf("got: %v\nwant: %v", pixels, expect)
	}
}

func TestGraph_GetPixelDatesFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphGetPixelDatesInput{ID: String(graphID), From: String("20180101"), To: String("20181231")}
	result, err := client.Graph().GetPixelDates(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraph_GetPixelDatesError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphGetPixelDatesInput{ID: String(graphID), From: String("20180101"), To: String("20181231")}
	_, err := client.Graph().GetPixelDates(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateStopwatchRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphStopwatchInput{ID: String(graphID)}
	param, err := client.Graph().createStopwatchRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPost {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPost)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/stopwatch", userName, graphID)
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

func TestGraph_Stopwatch(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphStopwatchInput{ID: String(graphID)}
	result, err := client.Graph().Stopwatch(input)

	testSuccess(t, result, err)
}

func TestGraph_StopwatchFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphStopwatchInput{ID: String(graphID)}
	result, err := client.Graph().Stopwatch(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_StopwatchError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphStopwatchInput{ID: String(graphID)}
	_, err := client.Graph().Stopwatch(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateGetRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphGetInput{ID: String(graphID)}
	param, err := client.Graph().createGetRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodGet {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodGet)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/graph-def", userName, graphID)
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

func TestGraph_Get(t *testing.T) {
	s := `{"id":"test-graph","name":"graph-name","unit":"commit","type":"int","color":"shibafu","timezone":"Asia/Tokyo","purgeCacheURLs":["https://camo.githubusercontent.com/xxx/xxxx"],"selfSufficient":"increment","isSecret":true,"publishOptionalData":true}`
	b := []byte(s)
	clientMock = &httpClientMock{statusCode: http.StatusOK, body: b}

	client := New(userName, token)
	input := &GraphGetInput{ID: String(graphID)}
	definition, err := client.Graph().Get(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	expect := &GraphDefinition{
		ID:                  "test-graph",
		Name:                "graph-name",
		Unit:                "commit",
		Type:                "int",
		Color:               "shibafu",
		TimeZone:            "Asia/Tokyo",
		PurgeCacheURLs:      []string{"https://camo.githubusercontent.com/xxx/xxxx"},
		SelfSufficient:      "increment",
		IsSecret:            true,
		PublishOptionalData: true,
		Result:              Result{IsSuccess: true, StatusCode: http.StatusOK},
	}
	if reflect.DeepEqual(definition, expect) == false {
		t.Errorf("got: %v\nwant: %v", definition, expect)
	}
}

func TestGraph_GetFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphGetInput{ID: String(graphID)}
	result, err := client.Graph().Get(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", result)
	}

	testAPIFailedResult(t, &result.Result, err)
}

func TestGraph_GetError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphGetInput{ID: String(graphID)}
	_, err := client.Graph().Get(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateAddRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphAddInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	param, err := client.Graph().createAddRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/add", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"quantity":"1"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraph_Add(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphAddInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	result, err := client.Graph().Add(input)

	testSuccess(t, result, err)
}

func TestGraph_AddFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphAddInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	result, err := client.Graph().Add(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_AddError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphAddInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	_, err := client.Graph().Add(input)

	testPageNotFoundError(t, err)
}

func TestGraph_CreateSubtractRequestParameter(t *testing.T) {
	client := New(userName, token)
	input := &GraphSubtractInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	param, err := client.Graph().createSubtractRequestParameter(input)
	if err != nil {
		t.Errorf("got: %v\nwant: nil", err)
	}

	if param.Method != http.MethodPut {
		t.Errorf("request method: %s\nwant: %s", param.Method, http.MethodPut)
	}

	expect := fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/subtract", userName, graphID)
	if param.URL != expect {
		t.Errorf("URL: %s\nwant: %s", param.URL, expect)
	}

	if param.Header[userToken] != token {
		t.Errorf("%s: %s\nwant: %s", userToken, param.Header[userToken], token)
	}

	s := `{"quantity":"1"}`
	b := []byte(s)
	if bytes.Equal(param.Body, b) == false {
		t.Errorf("Body: %s\nwant: %s", string(param.Body), s)
	}
}

func TestGraph_Subtract(t *testing.T) {
	clientMock = newOKMock()

	client := New(userName, token)
	input := &GraphSubtractInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	result, err := client.Graph().Subtract(input)

	testSuccess(t, result, err)
}

func TestGraph_SubtractFail(t *testing.T) {
	clientMock = newAPIFailedMock()

	client := New(userName, token)
	input := &GraphSubtractInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	result, err := client.Graph().Subtract(input)

	testAPIFailedResult(t, result, err)
}

func TestGraph_SubtractError(t *testing.T) {
	clientMock = newPageNotFoundMock()

	client := New(userName, token)
	input := &GraphSubtractInput{
		ID:       String(graphID),
		Quantity: String("1"),
	}
	_, err := client.Graph().Subtract(input)

	testPageNotFoundError(t, err)
}
