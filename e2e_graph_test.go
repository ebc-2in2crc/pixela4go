package pixela

import (
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
	}
	result, err := e2eClient.Graph().Create(input)
	if err != nil {
		t.Errorf("Graph.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Create() got: %+v\nwant: true", result)
	}
}

func testE2EGraphUpdate(t *testing.T) {
	input := &GraphUpdateInput{
		ID:   String(graphID),
		Unit: String("times"), // fix typo: "time" => "times"
	}
	result, err := e2eClient.Graph().Update(input)
	if err != nil {
		t.Errorf("Graph.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Graph.Update() got: %+v\nwant: true", result)
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
	expected := &GraphDefinition{
		ID:             graphID,
		Name:           "graph-name",
		Unit:           "times",
		Type:           GraphTypeInt,
		Color:          GraphColorShibafu,
		TimeZone:       "Asia/Tokyo",
		SelfSufficient: GraphSelfSufficientIncrement,
		Result: Result{
			IsSuccess: true,
			Message:   "",
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
	pixels, ok := result.Pixels.([]string)
	if ok == false {
		t.Errorf("Graph.GetPixelDates() got: %T\nwant: []string", result.Pixels)
	}
	if len(pixels) != 1 {
		t.Errorf("Graph.GetPixelDates() got: %+v\nwant: 1", len(pixels))
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
	if result.TotalPixelsCount != 1 {
		t.Errorf("Graph.Stats().TotalPixelsCount got: %+v\nwant: 1", result.TotalPixelsCount)
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
}
