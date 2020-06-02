package pixela

import (
	"reflect"
	"testing"
)

func testE2ENotificationCreate(t *testing.T) {
	input := &NotificationCreateInput{
		GraphID:   String(graphID),
		ID:        String("notification-id"),
		Name:      String("notification-name"),
		Target:    String(NotificationTargetQuantity),
		Condition: String(NotificationConditionGreaterThan),
		Threshold: String("3"),
		RemindBy:  String("23"),
		ChannelID: String("channel-id"),
	}
	result, err := e2eClient.Notification().Create(input)
	if err != nil {
		t.Errorf("Notification.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Notification.Create() got: %+v\nwant: true", result)
	}
}

func testE2ENotificationUpdate(t *testing.T) {
	input := &NotificationUpdateInput{
		GraphID: String(graphID),
		ID:      String("notification-id"),
		Name:    String("notification-new-name"),
	}
	result, err := e2eClient.Notification().Update(input)
	if err != nil {
		t.Errorf("Notification.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Notification.Update() got: %+v\nwant: true", result)
	}
}

func testE2ENotificationGetAll(t *testing.T) {
	input := &NotificationGetAllInput{GraphID: String(graphID)}
	result, err := e2eClient.Notification().GetAll(input)
	if err != nil {
		t.Errorf("Notification.GetAll() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Notification.GetAll() got: %+v\nwant: true", result)
	}
	expected := []NotificationDefinition{
		{
			ID:        "notification-id",
			Name:      "notification-new-name",
			Target:    NotificationTargetQuantity,
			Condition: NotificationConditionGreaterThan,
			Threshold: "3",
			RemindBy:  "23",
			ChannelID: "channel-id",
		},
	}
	if reflect.DeepEqual(result.Notifications, expected) == false {
		t.Errorf("Notification.GetAll() got: %+v\nwant: %+v", result.Notifications, expected)
	}
}

func testE2ENotificationDelete(t *testing.T) {
	input := &NotificationDeleteInput{
		GraphID: String(graphID),
		ID:      String("notification-id"),
	}
	result, err := e2eClient.Notification().Delete(input)
	if err != nil {
		t.Errorf("Notification.Delete() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Notification.Delete() got: %+v\nwant: true", result)
	}
}
