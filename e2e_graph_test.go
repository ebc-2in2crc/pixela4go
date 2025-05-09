package pixela

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func testE2EGraphCreate(t *testing.T) {
	input := &GraphCreateInput{
		ID:             String(graphID),
		Name:           String("graph-name"),
		Unit:           String("time"), // typo
		Type:           String(GraphTypeInt),
		Color:          String(GraphColorShibafu),
		TimeZone:       String("Asia/Tokyo"),
		SelfSufficient: String(GraphSelfSufficientIncrement),
		StartOnMonday:  Bool(true),
	}
	result, err := e2eClient.Graph().Create(input)
	if err != nil {
		t.Errorf("Graph.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Create() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.Create() got: %+v\nwant: %d", result, http.StatusOK)
	}
}

func testE2EGraphUpdate(t *testing.T) {
	input := &GraphUpdateInput{
		ID:            String(graphID),
		Unit:          String("times"), // fix typo: "time" => "times"
		StartOnMonday: Bool(true),
	}
	result, err := e2eClient.Graph().Update(input)
	if err != nil {
		t.Errorf("Graph.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Update() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.Update() got: %+v\nwant: %d", result, http.StatusOK)
	}
}

func testE2EGraphGetAll(t *testing.T) {
	result, err := e2eClient.Graph().GetAll()
	if err != nil {
		t.Errorf("Graph.GetAll() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.GetAll() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.GetAll() got: %+v\nwant: %d", result, http.StatusOK)
	}
	expected := []GraphDefinition{
		{
			ID:             graphID,
			Name:           "graph-name",
			Unit:           "times",
			Type:           GraphTypeInt,
			Color:          GraphColorShibafu,
			TimeZone:       "Asia/Tokyo",
			SelfSufficient: GraphSelfSufficientIncrement,
		},
	}
	if reflect.DeepEqual(result.Graphs, expected) == false {
		t.Errorf("Graph.GetAll() got: %+v\nwant: %+v", result.Graphs, expected)
	}
}

func testE2EGraphGet(t *testing.T) {
	input := &GraphGetInput{ID: String(graphID)}
	result, err := e2eClient.Graph().Get(input)
	if err != nil {
		t.Errorf("Graph.Get() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Get() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.Get() got: %+v\nwant: %d", result, http.StatusOK)
	}
	expected := &GraphDefinition{
		ID:             graphID,
		Name:           "graph-name",
		Unit:           "times",
		Type:           GraphTypeInt,
		Color:          GraphColorShibafu,
		TimeZone:       "Asia/Tokyo",
		SelfSufficient: GraphSelfSufficientIncrement,
		Result: Result{
			IsSuccess:  true,
			Message:    "",
			StatusCode: http.StatusOK,
		},
	}
	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Graph.Get() got: %+v\nwant: %+v", result, expected)
	}
}

func testE2EGraphStopwatch(t *testing.T) {
	// Start the measurement of the time.
	invokeGraphStopwatch(t)
	// End the measurement of the time.
	invokeGraphStopwatch(t)
}

func invokeGraphStopwatch(t *testing.T) {
	input := &GraphStopwatchInput{
		ID: String(graphID),
	}
	result, err := e2eClient.Graph().Stopwatch(input)
	if err != nil {
		t.Errorf("Graph.Stopwatch() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Stopwatch() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.Stopwatch() got: %+v\nwant: %d", result, http.StatusOK)
	}
}

func testE2EGraphGetPixelDates(t *testing.T) {
	input := &GraphGetPixelDatesInput{
		ID: String(graphID),
	}
	result, err := e2eClient.Graph().GetPixelDates(input)
	if err != nil {
		t.Errorf("Graph.GetPixelDates() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.GetPixelDates() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.GetPixelDates() got: %+v\nwant: %d", result, http.StatusOK)
	}
	pixels, ok := result.Pixels.([]string)
	if ok == false {
		t.Errorf("Graph.GetPixelDates() got: %T\nwant: []string", result.Pixels)
	}
	if len(pixels) != 1 {
		t.Errorf("Graph.GetPixelDates() got: %+v\nwant: 1", len(pixels))
	}
}

func testE2EGraphGetToday(t *testing.T) {
	// Test with ReturnEmpty = nil (default behavior)
	input1 := &GraphGetTodayInput{
		ID: String(graphID),
	}
	result1, err := e2eClient.Graph().GetToday(input1)
	if err != nil {
		t.Errorf("Graph.GetToday() got: %+v\nwant: nil", err)
	}
	if result1.IsSuccess == false {
		t.Errorf("Graph.GetToday() got: %+v\nwant: true", result1)
	}
	if result1.StatusCode != http.StatusOK {
		t.Errorf("Graph.GetToday() got: %+v\nwant: %d", result1, http.StatusOK)
	}

	// Test with ReturnEmpty = true
	input2 := &GraphGetTodayInput{
		ID:          String(graphID),
		ReturnEmpty: Bool(true),
	}
	result2, err := e2eClient.Graph().GetToday(input2)
	if err != nil {
		t.Errorf("Graph.GetToday() with returnEmpty=true got: %+v\nwant: nil", err)
	}
	if result2.IsSuccess == false {
		t.Errorf("Graph.GetToday() with returnEmpty=true got: %+v\nwant: true", result2)
	}
	if result2.StatusCode != http.StatusOK {
		t.Errorf("Graph.GetToday() with returnEmpty=true got: %+v\nwant: %d", result2, http.StatusOK)
	}
}

func testE2EGraphGetSVG(t *testing.T) {
	input := &GraphGetSVGInput{ID: String(graphID)}
	result, err := e2eClient.Graph().GetSVG(input)
	if err != nil {
		t.Errorf("Graph.GetSVG() got: %+v\nwant: nil", err)
	}
	if strings.Contains(result, "<svg xmlns") == false {
		t.Errorf("Graph.GetSVG() got: %+v\nwant: <svg xmlns", result)
	}
}

func testE2EGraphStats(t *testing.T) {
	input := &GraphStatsInput{ID: String(graphID)}
	result, err := e2eClient.Graph().Stats(input)
	if err != nil {
		t.Errorf("Graph.Stats() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Stats() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.Stats() got: %+v\nwant: %d", result, http.StatusOK)
	}
	if result.TotalPixelsCount != 1 {
		t.Errorf("Graph.Stats().TotalPixelsCount got: %+v\nwant: 1", result.TotalPixelsCount)
	}
}

func testE2EGraphUpdatePixels(t *testing.T) {
	input := &GraphUpdatePixelsInput{
		ID: String(graphID),
		Pixels: []PixelInput{
			{
				Date:     String("20180101"),
				Quantity: String("1"),
			},
			{
				Date:     String("20180102"),
				Quantity: String("2"),
			},
		},
	}
	result, err := e2eClient.Graph().UpdatePixels(input)
	if err != nil {
		t.Errorf("Graph.UpdatePixels() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.UpdatePixels() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.UpdatePixels() got: %+v\nwant: %d", result, http.StatusOK)
	}
}

func testE2EGraphDelete(t *testing.T) {
	input := &GraphDeleteInput{ID: String(graphID)}
	result, err := e2eClient.Graph().Delete(input)
	if err != nil {
		t.Errorf("Graph.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Create() got: %+v\nwant: true", result)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Graph.Delete() got: %+v\nwant: %d", result, http.StatusOK)
	}
}
