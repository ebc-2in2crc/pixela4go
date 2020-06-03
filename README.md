[English](README.md) | [日本語](README_ja.md)

# pixela4go

[![Build Status](https://travis-ci.com/ebc-2in2crc/pixela4go.svg?branch=master)](https://travis-ci.com/ebc-2in2crc/pixela4go)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/pixela4go?status.svg)](https://godoc.org/github.com/ebc-2in2crc/pixela4go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/pixela4go)](https://goreportcard.com/report/github.com/ebc-2in2crc/pixela4go)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/pixela4go.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/pixela4go.svg?label=version)

[Pixela](https://pixe.la/) API client for Go.

[![Cloning count](https://pixe.la/v1/users/ebc-2in2crc/graphs/pixela4go-clone)](https://pixe.la/v1/users/ebc-2in2crc/graphs/pixela4go-clone.html)

## Documentation

https://godoc.org/github.com/ebc-2in2crc/pixela4go

## Installation

```
$ go get -u github.com/ebc-2in2crc/pixela4go
```

## Usage

```go
package main

import (
	"log"
	
	"github.com/ebc-2in2crc/pixela4go"
)

func main() {
	client := pixela.New("YOUR_NAME", "YOUR_TOKEN")

	// Create new user
	uci := &pixela.UserCreateInput{
		AgreeTermsOfService: pixela.Bool(true),
		NotMinor:            pixela.Bool(true),
		ThanksCode:          pixela.String("thanks-code"),
	}
	result, err := client.User().Create(uci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// Create new slack channel
	cci := &pixela.ChannelCreateInput{
		ID:          pixela.String("channel-id"),
		Name:        pixela.String("channel-name"),
		Type:        pixela.String(pixela.ChannelTypeSlack),
		SlackDetail: &pixela4go.SlackDetail{
			URL:         pixela4go.String("https://hooks.slack.com/services/xxxx"),
			UserName:    pixela4go.String("slack-user-name"),
			ChannelName: pixela4go.String("slack-channel-name"),
		},
	}
	result, err = client.Channel().Create(cci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// Create new graph
	gci := &pixela.GraphCreateInput{
		ID:                  pixela.String("graph-id"),
		Name:                pixela.String("graph-name"),
		Unit:                pixela.String("commit"),
		Type:                pixela.String(pixela.GraphTypeInt),
		Color:               pixela.String(pixela.GraphColorShibafu),
		TimeZone:            pixela.String("Asia/Tokyo"),
		SelfSufficient:      pixela.String(pixela.GraphSelfSufficientIncrement),
		IsSecret:            pixela.Bool(true),
		PublishOptionalData: pixela.Bool(true),
	}
	result, err = client.Graph().Create(gci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// Register value
	pci := &pixela.PixelCreateInput{
		Date:         pixela.String("20180915"),
		Quantity:     pixela.String("5"),
		GraphID:      pixela.String("graph-id"),
	}
	result, err = client.Pixel().Create(pci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// Create new notification rule
	nci := &pixela.NotificationCreateInput{
		GraphID:   pixela.String("graph-id"),
		ID:        pixela.String("notification-id"),
		Name:      pixela.String("notification-name"),
		Target:    pixela.String(pixela.NotificationTargetQuantity),
		Condition: pixela.String(pixela.NotificationConditionGreaterThan),
		Threshold: pixela.String("3"),
		RemindBy:  pixela.String("23"),
		ChannelID: pixela.String("channel-id"),
	}
	result, err = client.Notification().Create(nci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// Create new webhook
	wci := &pixela.WebhookCreateInput{
		GraphID: pixela.String("graph-id"),
		Type:    pixela.String(pixela.WebhookTypeIncrement),
	}
	webhook, err := client.Webhook().Create(wci)
	if err != nil {
		log.Fatal(err)
	}
	if webhook.IsSuccess == false {
		log.Fatal(webhook.Message)
	}

	// Invoke webhook
	wii := &pixela.WebhookInvokeInput{WebhookHash: pixela.String("webhook-hash")}
	result, err = client.Webhook().Invoke(wii)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}
}
```

## Contribution

1. Fork this repository
2. Create your issue branch (`git checkout -b issue/:id`)
3. Change codes
4. Run test suite with the `make test` command and confirm that it passes
5. Run `make fmt`
6. Commit your changes (`git commit -am 'Add some feature'`)
7. Create new Pull Request

## Licence

[MIT](https://github.com/ebc-2in2crc/pixela4go/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
