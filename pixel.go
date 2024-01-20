package pixela

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Pixel manages communication with the Pixela pixel API.
type Pixel struct {
	UserName string
	Token    string
}

// Create records the quantity of the specified date as a "Pixel".
func (p *Pixel) Create(input *PixelCreateInput) (*Result, error) {
	return p.CreateWithContext(context.Background(), input)
}

// CreateWithContext records the quantity of the specified date as a "Pixel".
func (p *Pixel) CreateWithContext(ctx context.Context, input *PixelCreateInput) (*Result, error) {
	param, err := p.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel create parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// PixelCreateInput is input of Pixel.Create().
type PixelCreateInput struct {
	// GraphID is a required field
	GraphID *string `json:"-"`
	// Date is a required field
	Date *string `json:"date"`
	// Quantity is a required field
	Quantity     *string `json:"quantity"`
	OptionalData *string `json:"optionalData,omitempty"`
}

func (p *Pixel) createCreateRequestParameter(input *PixelCreateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	graphID := StringValue(input.GraphID)
	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s", p.UserName, graphID),
		Header: map[string]string{userToken: p.Token},
		Body:   b,
	}, nil
}

// Increment increments quantity "Pixel" of the day (it is used "timezone" setting if Graph's "timezone" is specified, if not specified, calculates it in "UTC").
// If the graph type is int then 1 added, and for float then 0.01 added.
func (p *Pixel) Increment(input *PixelIncrementInput) (*Result, error) {
	return p.IncrementWithContext(context.Background(), input)
}

// IncrementWithContext increments quantity "Pixel" of the day (it is used "timezone" setting if Graph's "timezone" is specified, if not specified, calculates it in "UTC").
// If the graph type is int then 1 added, and for float then 0.01 added.
func (p *Pixel) IncrementWithContext(ctx context.Context, input *PixelIncrementInput) (*Result, error) {
	param, err := p.createIncrementRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel increment parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// PixelIncrementInput is input of Pixel.Increment().
type PixelIncrementInput struct {
	// GraphID is a required field
	GraphID *string
}

func (p *Pixel) createIncrementRequestParameter(input *PixelIncrementInput) (*requestParameter, error) {
	graphID := StringValue(input.GraphID)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/increment", p.UserName, graphID),
		Header: map[string]string{contentLength: "0", userToken: p.Token},
		Body:   []byte{},
	}, nil
}

// Decrement decrements quantity "Pixel" of the day (it is used "timezone" setting if Graph's "timezone" is specified, if not specified, calculates it in "UTC").
// If the graph type is int then -1 added, and for float then -0.01 added.
func (p *Pixel) Decrement(input *PixelDecrementInput) (*Result, error) {
	return p.DecrementWithContext(context.Background(), input)
}

// DecrementWithContext decrements quantity "Pixel" of the day (it is used "timezone" setting if Graph's "timezone" is specified, if not specified, calculates it in "UTC").
// If the graph type is int then -1 added, and for float then -0.01 added.
func (p *Pixel) DecrementWithContext(ctx context.Context, input *PixelDecrementInput) (*Result, error) {
	param, err := p.createDecrementRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel decrement parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// PixelDecrementInput is input of Pixel.Decrement().
type PixelDecrementInput struct {
	// GraphID is a required field
	GraphID *string
}

func (p *Pixel) createDecrementRequestParameter(input *PixelDecrementInput) (*requestParameter, error) {
	graphID := StringValue(input.GraphID)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/decrement", p.UserName, graphID),
		Header: map[string]string{contentLength: "0", userToken: p.Token},
		Body:   []byte{},
	}, nil
}

// Get gets registered quantity as "Pixel".
func (p *Pixel) Get(input *PixelGetInput) (*Quantity, error) {
	return p.GetWithContext(context.Background(), input)
}

// GetWithContext gets registered quantity as "Pixel".
func (p *Pixel) GetWithContext(ctx context.Context, input *PixelGetInput) (*Quantity, error) {
	param, err := p.createGetRequestParameter(input)
	if err != nil {
		return &Quantity{}, errors.Wrapf(err, "failed to create pixel get parameter")
	}

	b, status, err := doRequest(ctx, param)
	if err != nil {
		return &Quantity{}, errors.Wrapf(err, "failed to do request")
	}

	var quantity Quantity
	quantity.StatusCode = status
	if err := json.Unmarshal(b, &quantity); err != nil {
		return &Quantity{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	quantity.IsSuccess = quantity.Message == ""
	return &quantity, nil
}

// PixelGetInput is input of Pixel.Get().
type PixelGetInput struct {
	// GraphID is a required field
	GraphID *string
	// Date is required field.
	Date *string
}

func (p *Pixel) createGetRequestParameter(input *PixelGetInput) (*requestParameter, error) {
	graphID := StringValue(input.GraphID)
	date := StringValue(input.Date)
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/%s", p.UserName, graphID, date),
		Header: map[string]string{userToken: p.Token},
		Body:   []byte{},
	}, nil
}

// Quantity ... registered quantity.
type Quantity struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData"`
	Result
}

// Update updates the quantity already registered as a "Pixel".
func (p *Pixel) Update(input *PixelUpdateInput) (*Result, error) {
	return p.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext updates the quantity already registered as a "Pixel".
func (p *Pixel) UpdateWithContext(ctx context.Context, input *PixelUpdateInput) (*Result, error) {
	param, err := p.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel update parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// PixelUpdateInput is input of Pixel.Update().
type PixelUpdateInput struct {
	// GraphID is a required field
	GraphID *string `json:"-"`
	// Date is required field.
	Date         *string `json:"-"`
	Quantity     *string `json:"quantity,omitempty"`
	OptionalData *string `json:"optionalData,omitempty"`
}

func (p *Pixel) createUpdateRequestParameter(input *PixelUpdateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	graphID := StringValue(input.GraphID)
	date := StringValue(input.Date)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/%s", p.UserName, graphID, date),
		Header: map[string]string{userToken: p.Token},
		Body:   b,
	}, nil
}

// Delete deletes the registered "Pixel".
func (p *Pixel) Delete(input *PixelDeleteInput) (*Result, error) {
	return p.DeleteWithContext(context.Background(), input)
}

// DeleteWithContext deletes the registered "Pixel".
func (p *Pixel) DeleteWithContext(ctx context.Context, input *PixelDeleteInput) (*Result, error) {
	param, err := p.createDeleteRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create pixel delete parameter")
	}

	return doRequestAndParseResponse(ctx, param)
}

// PixelDeleteInput is input of Pixel.Delete().
type PixelDeleteInput struct {
	// GraphID is a required field
	GraphID *string
	// Date is a required field
	Date *string
}

func (p *Pixel) createDeleteRequestParameter(input *PixelDeleteInput) (*requestParameter, error) {
	graphID := StringValue(input.GraphID)
	date := StringValue(input.Date)
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURLForV1+"/users/%s/graphs/%s/%s", p.UserName, graphID, date),
		Header: map[string]string{userToken: p.Token},
		Body:   []byte{},
	}, nil
}
