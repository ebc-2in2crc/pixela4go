package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Specify the target to be notified.
const (
	NotificationTargetQuantity = "quantity"
)

// Specify the condition used to judge whether to notify or not.
const (
	NotificationConditionGreaterThan = ">"
	NotificationConditionEqual       = "="
	NotificationConditionLessThan    = "<"
	NotificationConditionMultipleOf  = "multipleOf"
)

// A Notification manages communication with the Pixela notification API.
type Notification struct {
	UserName string
	Token    string
}

// Create creates a new notification rule.
func (n *Notification) Create(input *NotificationCreateInput) (*Result, error) {
	param, err := n.createCreateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create notification create parameter")
	}

	return doRequestAndParseResponse(param)
}

// NotificationCreateInput is input of Notification.Delete().
type NotificationCreateInput struct {
	// ID is a required field
	ID *string `json:"id"`
	// Name is a required field
	Name *string `json:"name"`
	// Target is a required field
	Target *string `json:"target"`
	// Condition is a required field
	Condition *string `json:"condition"`
	// Threshold is a required field
	Threshold *string `json:"threshold"`
	RemindBy  *string `json:"remindBy,omitempty"`
	// ChannelID is a required field
	ChannelID *string `json:"channelID"`
	// GraphID is a required field
	GraphID *string `json:"-"`
}

func (n *Notification) createCreateRequestParameter(input *NotificationCreateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	graphID := StringValue(input.GraphID)
	return &requestParameter{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", n.UserName, graphID),
		Header: map[string]string{userToken: n.Token},
		Body:   b,
	}, nil
}

// GetAll get all predefined notifications.
func (n *Notification) GetAll(input *NotificationGetAllInput) (*NotificationDefinitions, error) {
	param, err := n.createGetRequestParameter(input)
	if err != nil {
		return &NotificationDefinitions{}, errors.Wrapf(err, "failed to create notification get parameter")
	}

	b, err := doRequest(param)
	if err != nil {
		return &NotificationDefinitions{}, errors.Wrapf(err, "failed to do request")
	}

	var definitions NotificationDefinitions
	if err := json.Unmarshal(b, &definitions); err != nil {
		return &NotificationDefinitions{}, errors.Wrapf(err, "failed to unmarshal json")
	}

	definitions.IsSuccess = definitions.Message == ""
	return &definitions, nil
}

// NotificationGetAllInput is input of Notification.GetAll().
type NotificationGetAllInput struct {
	// GraphID is a required field
	GraphID *string `json:"-"`
}

// NotificationDefinitions is notification list.
type NotificationDefinitions struct {
	Notifications []NotificationDefinition `json:"notifications"`
	Result
}

// NotificationDefinition is notification definition.
type NotificationDefinition struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Condition string `json:"condition"`
	Threshold string `json:"threshold"`
	RemindBy  string `json:"remindBy"`
	ChannelID string `json:"channelID"`
}

func (n *Notification) createGetRequestParameter(input *NotificationGetAllInput) (*requestParameter, error) {
	graphID := StringValue(input.GraphID)
	return &requestParameter{
		Method: http.MethodGet,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications", n.UserName, graphID),
		Header: map[string]string{userToken: n.Token},
		Body:   nil,
	}, nil
}

// Update updates predefined notification rule.
func (n *Notification) Update(input *NotificationUpdateInput) (*Result, error) {
	param, err := n.createUpdateRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create notification update parameter")
	}

	return doRequestAndParseResponse(param)
}

// NotificationUpdateInput is input of Notification.Update().
type NotificationUpdateInput struct {
	// ID is a required field
	ID *string `json:"-"`
	// Name is a required field
	Name *string `json:"name"`
	// Target is a required field
	Target *string `json:"target"`
	// Condition is a required field
	Condition *string `json:"condition"`
	// Threshold is a required field
	Threshold *string `json:"threshold"`
	RemindBy  *string `json:"remindBy,omitempty"`
	// ChannelID is a required field
	ChannelID *string `json:"channelID"`
	// GraphID is a required field
	GraphID *string `json:"-"`
}

func (n *Notification) createUpdateRequestParameter(input *NotificationUpdateInput) (*requestParameter, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return &requestParameter{}, errors.Wrap(err, "failed to marshal json")
	}

	graphID := StringValue(input.GraphID)
	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodPut,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", n.UserName, graphID, ID),
		Header: map[string]string{userToken: n.Token},
		Body:   b,
	}, nil
}

// Delete deletes predefined notification settings.
func (n *Notification) Delete(input *NotificationDeleteInput) (*Result, error) {
	param, err := n.createDeleteRequestParameter(input)
	if err != nil {
		return &Result{}, errors.Wrapf(err, "failed to create notification delete parameter")
	}

	return doRequestAndParseResponse(param)
}

// NotificationDeleteInput is input of Notification.Delete().
type NotificationDeleteInput struct {
	// ID is a required field
	ID *string `json:"id"`
	// GraphID is a required field
	GraphID *string `json:"-"`
}

func (n *Notification) createDeleteRequestParameter(input *NotificationDeleteInput) (*requestParameter, error) {
	graphID := StringValue(input.GraphID)
	ID := StringValue(input.ID)
	return &requestParameter{
		Method: http.MethodDelete,
		URL:    fmt.Sprintf(APIBaseURL+"/users/%s/graphs/%s/notifications/%s", n.UserName, graphID, ID),
		Header: map[string]string{userToken: n.Token},
		Body:   nil,
	}, nil
}
