package pixela

import (
	"reflect"
	"testing"
)

func testE2EChannelCreate(t *testing.T) {
	input := &ChannelCreateInput{
		ID:   String("channel-id"),
		Name: String("channel-name"),
		Type: String(ChannelTypeSlack),
		SlackDetail: &SlackDetail{
			URL:         String("https://hooks.slack.com/services/xxxxxx"),
			UserName:    String("slack-user"),
			ChannelName: String("slack-channel-name"),
		},
	}
	result, err := e2eClient.Channel().Create(input)
	if err != nil {
		t.Errorf("Channel.Create() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Channel.Create() got: %+v\nwant: true", result)
	}
}

func testE2EChannelUpdate(t *testing.T) {
	input := &ChannelUpdateInput{
		ID:   String("channel-id"),
		Name: String("channel-new-name"),
	}
	result, err := e2eClient.Channel().Update(input)
	if err != nil {
		t.Errorf("Channel.Update() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Channel.Update() got: %+v\nwant: true", result)
	}
}

func testE2EChannelGetAll(t *testing.T) {
	result, err := e2eClient.Channel().GetAll()
	if err != nil {
		t.Errorf("Channel.GetAll() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Channel.GetAll() got: %+v\nwant: true", result)
	}
	expected := []ChannelDefinition{
		{
			ID:   "channel-id",
			Name: "channel-new-name",
			Type: ChannelTypeSlack,
			Detail: SlackDetail{
				URL:         String("https://hooks.slack.com/services/xxxxxx"),
				UserName:    String("slack-user"),
				ChannelName: String("slack-channel-name"),
			},
		},
	}
	if reflect.DeepEqual(result.Channels, expected) == false {
		t.Errorf("Channel.GetAll() got: %+v\nwant: %+v", result.Channels, expected)
	}
}

func testE2EChannelDelete(t *testing.T) {
	input := &ChannelDeleteInput{ID: String("channel-id")}
	result, err := e2eClient.Channel().Delete(input)
	if err != nil {
		t.Errorf("Channel.Delete() got: %+v\nwant: nil", err)
	}
	if result.IsSuccess == false {
		t.Errorf("Channel.Delete() got: %+v\nwant: true", result)
	}
}
