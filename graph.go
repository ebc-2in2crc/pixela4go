package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Graph manages communication with the Pixela graph API.
type Graph struct {
	UserName string
	Token    string
}

// Create creates a new pixelation graph definition.
func (g *Graph) Create(input *GraphCreateInput) (*Result, error) {
	return g.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new pixelation graph definition.
func (g *Graph) CreateWithContext(ctx context.Context, input *GraphCreateInput) (*Result, error) {
	param, err := g.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph create parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphCreateInput is input of Graph.Create().
type GraphCreateInput struct {
	// ID is a required field
	ID *string `json:"id"`
	// Name is a required field
	Name *string `json:"name"`
	// Unit is a required field
	Unit *string `json:"unit"`
	// Type is a required field
	Type *string `json:"type"`
	// Color is a required field
	Color               *string `json:"color"`
	TimeZone            *string `json:"timezone,omitempty"`
	SelfSufficient      *string `json:"selfSufficient,omitempty"`
	IsSecret            *bool   `json:"isSecret,omitempty"`
	PublishOptionalData *bool   `json:"publishOptionalData,omitempty"`
}

func (g *Graph) createCreateRequestParameter(input *GraphCreateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs", g.UserName),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

// It is the type of quantity to be handled in the graph.
// Only int or float are supported.
const (
	GraphTypeInt   = "int"
	GraphTypeFloat = "float"
)

// Defines the display color of the pixel in the pixelation graph.
// shibafu (green), momiji (red), sora (blue), ichou (yellow), ajisai (purple) and kuro (black) are supported as color kind.
const (
	GraphColorShibafu = "shibafu"
	GraphColorMomiji  = "momiji"
	GraphColorSora    = "sora"
	GraphColorIchou   = "ichou"
	GraphColorAjisai  = "ajisai"
	GraphColorKuro    = "kuro"
)

// If SVG graph with this field increment or decrement is referenced, Pixel of this graph itself will be incremented or decremented.
// It is suitable when you want to record the PVs on a web page or site simultaneously.
// The specification of increment or decrement is the same as Increment a Pixel and Decrement a Pixel with webhook.
// If not specified, it is treated as none .
const (
	GraphSelfSufficientIncrement = "increment"
	GraphSelfSufficientDecrement = "decrement"
	GraphSelfSufficientNone      = "none"
)

// GetAll gets all predefined pixelation graph definitions.
func (g *Graph) GetAll() (*GraphDefinitions, error) {
	return g.GetAllWithContext(context.Background())
}

// GetAllWithContext gets all predefined pixelation graph definitions.
func (g *Graph) GetAllWithContext(ctx context.Context) (*GraphDefinitions, error) {
	param, err := g.createGetAllRequestParameter()
	if err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to create get all graph parameter")
	}

	b, status, err := doRequest(ctx, param)
	if err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var definitions GraphDefinitions
	definitions.StatusCode = status
	if err := json.Unmarshal(b, &definitions); err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definitions.IsSuccess = definitions.Message == ""
	return &definitions, nil
}

func (g *Graph) createGetAllRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs", g.UserName),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// GraphDefinitions is graph definition list.
type GraphDefinitions struct {
	Graphs []GraphDefinition `json:"graphs"`
	Result
}

// GraphDefinition is graph definition.
type GraphDefinition struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Unit                string   `json:"unit"`
	Type                string   `json:"type"`
	Color               string   `json:"color"`
	TimeZone            string   `json:"timezone"`
	PurgeCacheURLs      []string `json:"purgeCacheURLs"`
	SelfSufficient      string   `json:"selfSufficient"`
	IsSecret            bool     `json:"isSecret"`
	PublishOptionalData bool     `json:"publishOptionalData"`
	Result
}

// GetSVG get a graph expressed in SVG format diagram that based on the registered information.
func (g *Graph) GetSVG(input *GraphGetSVGInput) (string, error) {
	return g.GetSVGWithContext(context.Background(), input)
}

// GetSVGWithContext get a graph expressed in SVG format diagram that based on the registered information.
func (g *Graph) GetSVGWithContext(ctx context.Context, input *GraphGetSVGInput) (string, error) {
	param, err := g.createGetSVGRequestParameter(input)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create get svg parameter")
	}

	b, err := mustDoRequest(ctx, param)
	if err != nil {
		return "", errors.Wrapf(err, "failed to do request")
	}

	return string(b), nil
}

// GraphGetSVGInput is input of Graph.GetSVG().
type GraphGetSVGInput struct {
	// ID is a required field
	ID         *string `json:"-"`
	Date       *string `json:"date,omitempty"`
	Mode       *string `json:"mode,omitempty"`
	Appearance *string `json:"appearance,omitempty"`
}

func (g *Graph) createGetSVGRequestParameter(input *GraphGetSVGInput) (*requestParameter, error) {
	ID := StringValue(input.ID)
	date := StringValue(input.Date)
	if date != "" {
		date = "date=" + date
	}
	mode := StringValue(input.Mode)
	if mode != "" {
		mode = "mode=" + StringValue(input.Mode)
	}
	appearance := StringValue(input.Appearance)
	if appearance != "" {
		appearance = "appearance=" + StringValue(input.Appearance)
	}

	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s?%s&%s&%s", g.UserName, ID, date, mode, appearance),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// Specify the graph display mode.
// Supported modes are short (for displaying only about 90 days), badge (Badge format pasted on GitHub README.
// Information for the last 49 days is expressed in 7 pixels.), and line .
const (
	GraphModeShort = "short"
	GraphModeBadge = "badge"
	GraphModeLine  = "line"
)

// Specify the graph display mode in html format.
const (
	GraphModeSimple      = "simple"
	GraphModeSimpleShort = "simple-short"
)

// Dark theme
const (
	GraphAppearanceDark = "dark"
)

// UpdatePixels is used to register multiple Pixels (quantities for a specific day) at a time.
func (g *Graph) UpdatePixels(input *GraphUpdatePixelsInput) (*Result, error) {
	return g.UpdatePixelsWithContext(context.Background(), input)
}

// UpdatePixelsWithContext is used to register multiple Pixels (quantities for a specific day) at a time.
func (g *Graph) UpdatePixelsWithContext(ctx context.Context, input *GraphUpdatePixelsInput) (*Result, error) {
	param, err := g.createUpdatePixelsRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph update pixels parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphUpdatePixelsInput is input of Graph.UpdatePixels().
type GraphUpdatePixelsInput struct {
	// ID is a required field
	ID     *string      `json:"-"`
	Pixels []PixelInput `json:"-"`
}

// PixelInput is input of Graph.UpdatePixels().
type PixelInput struct {
	// Date is a required field
	Date *string `json:"date"`
	// Quantity is a required field
	Quantity     *string `json:"quantity"`
	OptionalData *string `json:"optionalData,omitempty"`
}

func (g *Graph) createUpdatePixelsRequestParameter(input *GraphUpdatePixelsInput) (*requestParameter, error) {
	b, err := json.Marshal(input.Pixels)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/pixels", g.UserName, ID),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

// URL displays the details of the graph in html format.
func (g *Graph) URL(input *GraphURLInput) string {
	ID := StringValue(input.ID)
	mode := StringValue(input.Mode)
	if mode == "" {
		return fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s.html", g.UserName, ID)
	}

	return fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s.html?mode=%s", g.UserName, ID, mode)
}

// GraphURLInput is input of Graph.GetURL().
type GraphURLInput struct {
	// ID is a required field
	ID   *string `json:"-"`
	Mode *string
}

// Stats is various statistics based on the registered information.
type Stats struct {
	TotalPixelsCount  int     `json:"totalPixelsCount"`
	MaxQuantity       int     `json:"maxQuantity"`
	MaxDate           string  `json:"maxDate"`
	MinQuantity       int     `json:"minQuantity"`
	MinDate           string  `json:"minDate"`
	TotalQuantity     int     `json:"totalQuantity"`
	AvgQuantity       float64 `json:"avgQuantity"`
	TodaysQuantity    int     `json:"todaysQuantity"`
	YesterdayQuantity int     `json:"yesterdayQuantity"`
	Result
}

// Stats gets various statistics based on the registered information.
func (g *Graph) Stats(input *GraphStatsInput) (*Stats, error) {
	return g.StatsWithContext(context.Background(), input)
}

// StatsWithContext gets various statistics based on the registered information.
func (g *Graph) StatsWithContext(ctx context.Context, input *GraphStatsInput) (*Stats, error) {
	param, err := g.createStatsRequestParameter(input)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create graph stats request parameter")
	}

	b, status, err := doRequest(ctx, param)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to do request")
	}

	var stats Stats
	stats.StatusCode = status
	if err := json.Unmarshal(b, &stats); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal json")
	}

	stats.IsSuccess = stats.Message == ""
	return &stats, nil
}

// GraphStatsInput is input of Graph.Stats().
type GraphStatsInput struct {
	// ID is a required field
	ID *string
}

func (g *Graph) createStatsRequestParameter(input *GraphStatsInput) (*requestParameter, error) {
	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/stats", g.UserName, ID),
		Header: map[string]string{},
		Body:   []byte{},
	}, nil
}

// Update updates predefined pixelation graph definitions.
// The items that can be updated are limited as compared with the pixelation graph definition creation.
func (g *Graph) Update(input *GraphUpdateInput) (*Result, error) {
	return g.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext updates predefined pixelation graph definitions.
// The items that can be updated are limited as compared with the pixelation graph definition creation.
func (g *Graph) UpdateWithContext(ctx context.Context, input *GraphUpdateInput) (*Result, error) {
	param, err := g.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph update parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphUpdateInput is input of Graph.Update().
type GraphUpdateInput struct {
	// ID is a required field
	ID                  *string  `json:"-"`
	Name                *string  `json:"name,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
	Color               *string  `json:"color,omitempty"`
	TimeZone            *string  `json:"timezone,omitempty"`
	PurgeCacheURLs      []string `json:"purgeCacheURLs,omitempty"`
	SelfSufficient      *string  `json:"selfSufficient,omitempty"`
	IsSecret            *bool    `json:"isSecret,omitempty"`
	PublishOptionalData *bool    `json:"publishOptionalData,omitempty"`
}

func (g *Graph) createUpdateRequestParameter(input *GraphUpdateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", g.UserName, ID),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

// Delete deletes the predefined pixelation graph definition.
func (g *Graph) Delete(input *GraphDeleteInput) (*Result, error) {
	return g.DeleteWithContext(context.Background(), input)
}

// DeleteWithContext deletes the predefined pixelation graph definition.
func (g *Graph) DeleteWithContext(ctx context.Context, input *GraphDeleteInput) (*Result, error) {
	param, err := g.createDeleteRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph delete parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphDeleteInput is input of Graph.Delete().
type GraphDeleteInput struct {
	// ID is a required field
	ID *string `json:"-"`
}

func (g *Graph) createDeleteRequestParameter(input *GraphDeleteInput) (*requestParameter, error) {
	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", g.UserName, ID),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// GetPixelDates gets a Date list of Pixel registered in the graph specified by graphID.
// You can specify a period with from and to parameters.
//
// If you do not specify both from and to;
// You will get a list of 365 days ago from today.
//
// If you specify from only;
// You will get a list of 365 days from from date.
//
// If you specify to only;
// You will get a list of 365 days ago from to date.
//
// If you specify both from andto;
// You will get a list you specify.
// You can not specify a period greater than 365 days.
func (g *Graph) GetPixelDates(input *GraphGetPixelDatesInput) (*Pixels, error) {
	return g.GetPixelDatesWithContext(context.Background(), input)
}

// GetPixelDatesWithContext gets a Date list of Pixel registered in the graph specified by graphID.
// You can specify a period with from and to parameters.
//
// If you do not specify both from and to;
// You will get a list of 365 days ago from today.
//
// If you specify from only;
// You will get a list of 365 days from from date.
//
// If you specify to only;
// You will get a list of 365 days ago from to date.
//
// If you specify both from andto;
// You will get a list you specify.
// You can not specify a period greater than 365 days.
func (g *Graph) GetPixelDatesWithContext(ctx context.Context, input *GraphGetPixelDatesInput) (*Pixels, error) {
	param, err := g.createGetPixelDatesRequestParameter(input)
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to create get pixel dates parameter")
	}

	b, status, err := doRequest(ctx, param)
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to do request")
	}

	pixels, err := unmarshalPixels(b, BoolValue(input.WithBody))
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	pixels.StatusCode = status
	pixels.IsSuccess = pixels.Message == ""
	return pixels, nil
}

// GraphGetPixelDatesInput is input of Graph.GetPixelDates().
type GraphGetPixelDatesInput struct {
	// ID is a required field
	ID       *string
	From     *string
	To       *string
	WithBody *bool
}

// Pixels is Date list of Pixel registered in the graph.
type Pixels struct {
	// Pixels as []PixelWithBody when `withBody` is true.
	// Pixels as []string when `withBody` is false.
	Pixels interface{}
	Result
}

// PixelWithBody is Date of Pixel registered in the graph.
type PixelWithBody struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
}

type pixelsWithBody struct {
	Pixels []PixelWithBody `json:"pixels"`
	Result
}

type pixelsWithNoBody struct {
	Pixels []string `json:"pixels"`
	Result
}

func unmarshalPixels(b []byte, withBody bool) (*Pixels, error) {
	if withBody {
		return unmarshalPixelsWithBody(b)
	}
	return unmarshalPixelsNoBody(b)
}

func unmarshalPixelsWithBody(b []byte) (*Pixels, error) {
	var pixels pixelsWithBody
	if err := json.Unmarshal(b, &pixels); err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	return &Pixels{
		Pixels: pixels.Pixels,
		Result: pixels.Result,
	}, nil
}

func unmarshalPixelsNoBody(b []byte) (*Pixels, error) {
	var pixels pixelsWithNoBody
	if err := json.Unmarshal(b, &pixels); err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	return &Pixels{
		Pixels: pixels.Pixels,
		Result: pixels.Result,
	}, nil
}

func (g *Graph) createGetPixelDatesRequestParameter(input *GraphGetPixelDatesInput) (*requestParameter, error) {
	ID := StringValue(input.ID)
	from := StringValue(input.From)
	if from != "" {
		from = "from=" + from
	}
	to := StringValue(input.To)
	if to != "" {
		to = "to=" + to
	}
	withBody := ""
	if BoolValue(input.WithBody) {
		withBody = "withBody=true"
	}
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/pixels?%s&%s&%s", g.UserName, ID, from, to, withBody),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// Stopwatch start and end the measurement of the time.
func (g *Graph) Stopwatch(input *GraphStopwatchInput) (*Result, error) {
	return g.StopwatchWithContext(context.Background(), input)
}

// StopwatchWithContext start and end the measurement of the time.
func (g *Graph) StopwatchWithContext(ctx context.Context, input *GraphStopwatchInput) (*Result, error) {
	param, err := g.createStopwatchRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph stopwatch parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphStopwatchInput is input of Graph.Stopwatch().
type GraphStopwatchInput struct {
	// ID is a required field
	ID *string
}

func (g *Graph) createStopwatchRequestParameter(input *GraphStopwatchInput) (*requestParameter, error) {
	graphID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/stopwatch", g.UserName, graphID),
		Header: map[string]string{contentLength: "0", userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// Get gets predefined pixelation graph definitions.
func (g *Graph) Get(input *GraphGetInput) (*GraphDefinition, error) {
	return g.GetWithContext(context.Background(), input)
}

// GetWithContext gets predefined pixelation graph definitions.
func (g *Graph) GetWithContext(ctx context.Context, input *GraphGetInput) (*GraphDefinition, error) {
	param, err := g.createGetRequestParameter(input)
	if err != nil {
		return &GraphDefinition{}, errors.Wrapf(err, "failed to create get graph parameter")
	}

	b, status, err := doRequest(ctx, param)
	if err != nil {
		return &GraphDefinition{}, errors.Wrapf(err, "failed to do request")
	}

	var definition GraphDefinition
	definition.StatusCode = status
	if err := json.Unmarshal(b, &definition); err != nil {
		return &GraphDefinition{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definition.IsSuccess = definition.Message == ""
	return &definition, nil
}

func (g *Graph) createGetRequestParameter(input *GraphGetInput) (*requestParameter, error) {
	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/graph-def", g.UserName, ID),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// GraphGetInput is input of Graph.Get().
type GraphGetInput struct {
	// ID is a required field
	ID *string `json:"-"`
}

// Add quantity to the "Pixel" of the day.
func (g *Graph) Add(input *GraphAddInput) (*Result, error) {
	return g.AddWithContext(context.Background(), input)
}

// AddWithContext quantity to the "Pixel" of the day.
func (g *Graph) AddWithContext(ctx context.Context, input *GraphAddInput) (*Result, error) {
	param, err := g.createAddRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph add parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphAddInput is input of Graph.Add().
type GraphAddInput struct {
	// ID is a required field
	ID *string `json:"-"`
	// Quantity is a required field
	Quantity *string `json:"quantity"`
}

func (g *Graph) createAddRequestParameter(input *GraphAddInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	graphID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/add", g.UserName, graphID),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

// Subtract quantity from the "Pixel" of the day.
func (g *Graph) Subtract(input *GraphSubtractInput) (*Result, error) {
	return g.SubtractWithContext(context.Background(), input)
}

// SubtractWithContext quantity from the "Pixel" of the day.
func (g *Graph) SubtractWithContext(ctx context.Context, input *GraphSubtractInput) (*Result, error) {
	param, err := g.createSubtractRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph add parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// GraphSubtractInput is input of Graph.Subtract().
type GraphSubtractInput struct {
	// ID is a required field
	ID *string `json:"-"`
	// Quantity is a required field
	Quantity *string `json:"quantity"`
}

func (g *Graph) createSubtractRequestParameter(input *GraphSubtractInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	graphID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/subtract", g.UserName, graphID),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}
