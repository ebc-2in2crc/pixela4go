package pixela

import (
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
	param, err := g.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph create parameter")
	}

	return doRequestAndParseResponse(param)
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
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs", g.UserName),
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
	GraphColorMomiji  = "momiji "
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
	param, err := g.createGetAllRequestParameter()
	if err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to create get all graph parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var definitions GraphDefinitions
	if err := json.Unmarshal(b, &definitions); err != nil {
		return &GraphDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definitions.IsSuccess = definitions.Message == ""
	return &definitions, nil
}

func (g *Graph) createGetAllRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs", g.UserName),
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
}

// GetSVG get a graph expressed in SVG format diagram that based on the registered information.
func (g *Graph) GetSVG(input *GraphGetSVGInput) (string, error) {
	param, err := g.createGetSVGRequestParameter(input)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create get svg parameter")
	}

	b, err := mustDoRequest(param)
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
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s?%s&%s&%s", g.UserName, ID, date, mode, appearance),
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

// URL displays the details of the graph in html format.
func (g *Graph) URL(input *GraphURLInput) string {
	ID := StringValue(input.ID)
	mode := StringValue(input.Mode)
	if mode == "" {
		return fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s.html", g.UserName, ID)
	}

	return fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s.html?mode=%s", g.UserName, ID, mode)
}

// GraphURLInput is input of Graph.GetURL().
type GraphURLInput struct {
	// ID is a required field
	ID   *string `json:"-"`
	Mode *string
}

// GraphsURL displays graph list by detail in html format.
func (g *Graph) GraphsURL() string {
	return fmt.Sprintf(APIBaseURL+"/users/%s/graphs.html", g.UserName)
}

// Stats is various statistics based on the registered information.
type Stats struct {
	TotalPixelsCount int     `json:"totalPixelsCount"`
	MaxQuantity      int     `json:"maxQuantity"`
	MinQuantity      int     `json:"minQuantity"`
	TotalQuantity    int     `json:"totalQuantity"`
	AvgQuantity      float64 `json:"avgQuantity"`
	TodaysQuantity   int     `json:"todaysQuantity"`
	Result
}

// Stats gets various statistics based on the registered information.
func (g *Graph) Stats(input *GraphStatsInput) (*Stats, error) {
	param, err := g.createStatsRequestParameter(input)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create graph stats request parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to do request")
	}

	var stats Stats
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
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/stats", g.UserName, ID),
		Header: map[string]string{},
		Body:   []byte{},
	}, nil
}

// Update updates predefined pixelation graph definitions.
// The items that can be updated are limited as compared with the pixelation graph definition creation.
func (g *Graph) Update(input *GraphUpdateInput) (*Result, error) {
	param, err := g.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph update parameter")
	}

	return doRequestAndParseResponse(param)
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
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", g.UserName, ID),
		Header: map[string]string{userToken: g.Token},
		Body:   b,
	}, nil
}

// Delete deletes the predefined pixelation graph definition.
func (g *Graph) Delete(input *GraphDeleteInput) (*Result, error) {
	param, err := g.createDeleteRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph delete parameter")
	}

	return doRequestAndParseResponse(param)
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
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s", g.UserName, ID),
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
	param, err := g.createGetPixelDatesRequestParameter(input)
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to create get pixel dates parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to do request")
	}

	var pixels Pixels
	if err := json.Unmarshal(b, &pixels); err != nil {
		return &Pixels{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	pixels.IsSuccess = pixels.Message == ""
	return &pixels, nil
}

// GraphGetPixelDatesInput is input of Graph.GetPixelDates().
type GraphGetPixelDatesInput struct {
	// ID is a required field
	ID   *string
	From *string
	To   *string
}

// Pixels is Date list of Pixel registered in the graph.
type Pixels struct {
	Pixels []string `json:"pixels"`
	Result
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
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/pixels?%s&%s", g.UserName, ID, from, to),
		Header: map[string]string{userToken: g.Token},
		Body:   []byte{},
	}, nil
}

// Stopwatch start and end the measurement of the time.
func (g *Graph) Stopwatch(input *GraphStopwatchInput) (*Result, error) {
	param, err := g.createStopwatchRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create graph stopwatch parameter")
	}

	return doRequestAndParseResponse(param)
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
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/stopwatch", g.UserName, graphID),
		Header: map[string]string{contentLength: "0", userToken: g.Token},
		Body:   []byte{},
	}, nil
}
