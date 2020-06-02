package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// A Channel manages communication with the Pixela graph API.
type Channel struct {
	UserName string
	Token    string
}

// It is the channel type for notification. Onlyslack is supported.
const (
	ChannelTypeSlack = "slack"
)

// Create creates a new channel.
func (c *Channel) Create(input *ChannelCreateInput) (*Result, error) {
	param, err := c.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create channel create parameter")
	}

	return doRequestAndParseResponse(param)
}

// ChannelCreateInput is input of Channel.Create().
type ChannelCreateInput struct {
	// ID is a required field
	ID *string `json:"id"`
	// Name is a required field
	Name *string `json:"name"`
	// Type is a required field
	Type *string `json:"type"`
	// SlackDetail is a required field when Type is "slack".
	SlackDetail *SlackDetail `json:"detail,omitempty"`
}

// SlackDetail is channel detail settings when type is slack.
type SlackDetail struct {
	// URL is a required field
	URL *string `json:"url"`
	// UserName is a required field
	UserName *string `json:"userName"`
	// ChannelName is a required field
	ChannelName *string `json:"channelName"`
}

func (c *Channel) createCreateRequestParameter(input *ChannelCreateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels", c.UserName),
		Header: map[string]string{userToken: c.Token},
		Body:   b,
	}, nil
}

// GetAll gets all predefined channels.
func (c *Channel) GetAll() (*ChannelDefinitions, error) {
	param, err := c.createGetRequestParameter()
	if err != nil {
		return &ChannelDefinitions{}, errors.Wrapf(err, "failed to create get all channels parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &ChannelDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var raw rawChannelDefinitions
	if err := json.Unmarshal(b, &raw); err != nil {
		return &ChannelDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	var definitions ChannelDefinitions
	definitions.Channels = make([]ChannelDefinition, len(raw.Channels))
	for i, v := range raw.Channels {
		definitions.Channels[i], err = createChannelDefinition(v)
		if err != nil {
			return &ChannelDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
		}
	}

	definitions.Message = raw.Message
	definitions.IsSuccess = raw.Message == ""
	return &definitions, nil
}

func createChannelDefinition(raw rawChannelDefinition) (ChannelDefinition, error) {
	d := ChannelDefinition{
		ID:   raw.ID,
		Name: raw.Name,
		Type: raw.Type,
	}
	switch d.Type {
	case ChannelTypeSlack:
		var slack SlackDetail
		err := json.Unmarshal(raw.Detail, &slack)
		if err != nil {
			return ChannelDefinition{}, errors.Wrapf(err, "failed to unmarshal json")
		}
		d.Detail = slack
	default:
		return ChannelDefinition{}, errors.Errorf("unsupported type: %s", d.Type)
	}
	return d, nil
}

func (c *Channel) createGetRequestParameter() (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels", c.UserName),
		Header: map[string]string{userToken: c.Token},
	}, nil
}

type rawChannelDefinitions struct {
	Channels []rawChannelDefinition `json:"channels"`
	Result
}

type rawChannelDefinition struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`
	Detail json.RawMessage `json:"detail"`
}

// ChannelDefinitions is channel definition list.
type ChannelDefinitions struct {
	Channels []ChannelDefinition `json:"channels"`
	Result
}

// ChannelDefinition is channel definition.
type ChannelDefinition struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Detail interface{} `json:"detail"`
}

// Update updates a predefined channel.
func (c *Channel) Update(input *ChannelUpdateInput) (*Result, error) {
	param, err := c.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create channel update parameter")
	}

	return doRequestAndParseResponse(param)
}

// ChannelUpdateInput is input of Channel.Update().
type ChannelUpdateInput struct {
	// ID is a required field
	ID   *string `json:"-"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	// SlackDetail is a required field when Type is "slack".
	SlackDetail *SlackDetail `json:"detail,omitempty"`
}

func (c *Channel) createUpdateRequestParameter(input *ChannelUpdateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels/%s", c.UserName, StringValue(input.ID)),
		Header: map[string]string{userToken: c.Token},
		Body:   b,
	}, nil
}

// Delete deletes the predefined channel settings.
func (c *Channel) Delete(input *ChannelDeleteInput) (*Result, error) {
	param, err := c.createDeleteRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create channel delete parameter")
	}

	return doRequestAndParseResponse(param)
}

// ChannelDeleteInput is input of Channel.Delete().
type ChannelDeleteInput struct {
	// ID is a required field
	ID *string
}

func (c *Channel) createDeleteRequestParameter(input *ChannelDeleteInput) (*requestParameter, error) {
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/channels/%s", c.UserName, StringValue(input.ID)),
		Header: map[string]string{userToken: c.Token},
		Body:   nil,
	}, nil
}
