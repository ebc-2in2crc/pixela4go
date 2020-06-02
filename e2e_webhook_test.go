package pixela

import (
	"reflect"
	"testing"
)

var webhookHash string

func testE2EWebhookCreate(t *testing.T) {
	input := &WebhookCreateInput{
		GraphID: String(graphID),
		Type:    String(WebhookTypeIncrement),
	}
	result, err := e2eClient.Webhook().Create(input)
	if err != nil {
		t.Errorf("Webhook.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Webhook.Create() got: %+v\nwant: true", result)
	}

	webhookHash = result.WebhookHash
}

func testE2EWebhookInvoke(t *testing.T) {
	input := &WebhookInvokeInput{WebhookHash: String(webhookHash)}
	result, err := e2eClient.Webhook().Invoke(input)
	if err != nil {
		t.Errorf("Webhook.Invoke() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Webhook.Invoke() got: %+v\nwant: true", result)
	}
}

func testE2EWebhookGetAll(t *testing.T) {
	result, err := e2eClient.Webhook().GetAll()
	if err != nil {
		t.Errorf("Webhook.GetAll() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Webhook.GetAll() got: %+v\nwant: true", result)
	}
	expected := []WebhookDefinition{
		{
			GraphID:     graphID,
			Type:        WebhookTypeIncrement,
			WebhookHash: webhookHash,
		},
	}
	if reflect.DeepEqual(result.Webhooks, expected) == false {
		t.Errorf("Webhook.GetAll() got: %+v\nwant: %+v", result.Webhooks, expected)
	}
}

func testE2EWebhookDelete(t *testing.T) {
	input := &WebhookDeleteInput{WebhookHash: String(webhookHash)}
	result, err := e2eClient.Webhook().Delete(input)
	if err != nil {
		t.Errorf("Webhook.Delete() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Webhook.Delete() got: %+v\nwant: true", result)
	}
}
