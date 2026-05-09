[English](README.md) | [日本語](README_ja.md)

# pixela4go

![CI](https://github.com/ebc-2in2crc/pixela4go/actions/workflows/pr.yml/badge.svg)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/pixela4go?status.svg)](https://pkg.go.dev/github.com/ebc-2in2crc/pixela4go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/pixela4go)](https://goreportcard.com/report/github.com/ebc-2in2crc/pixela4go)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/pixela4go.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/pixela4go.svg?label=version)

Go 用の [Pixela](https://pixe.la/) API クライアントです。

[![Cloning count](https://pixe.la/v1/users/ebc-2in2crc/graphs/pixela4go-clone)](https://pixe.la/v1/users/ebc-2in2crc/graphs/pixela4go-clone.html)

## ドキュメント

https://godoc.org/github.com/ebc-2in2crc/pixela4go

## インストール

```
$ go get -u github.com/ebc-2in2crc/pixela4go
```

## 使い方

```go
package main

import (
	"context"
	"log"

	pixela "github.com/ebc-2in2crc/pixela4go"
)

func main() {
	// API 呼び出しが拒否されたときにリトライする場合、リトライ回数を指定します。
	// もしリトライしないなら指定する必要はありません。
	pixela.RetryCount = 10
	
	client := pixela.New("YOUR_NAME", "YOUR_TOKEN")

	// 新しいユーザーを作る
	uci := &pixela.UserCreateInput{
		AgreeTermsOfService: pixela.Bool(true),
		NotMinor:            pixela.Bool(true),
		ThanksCode:          pixela.String("thanks-code"),
	}
	result, err := client.User().CreateWithContext(context.Background(), uci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// ユーザープロフィールページを更新する
	upi := &pixela.UserProfileUpdateInput{
		DisplayName:       pixela.String("display-name"),
		GravatarIconEmail: pixela.String("gravatar-icon-email"),
		Title:             pixela.String("title"),
		Timezone:          pixela.String("Asia/Tokyo"),
		AboutURL:          pixela.String("https://github.com/ebc-2in2crc"),
		ContributeURLs:    []string{},
		PinnedGraphID:     pixela.String("pinned-graph-id"),
	}
	result, err = client.UserProfile().UpdateWithContext(context.Background(), upi)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 新しいグラフを作る
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
	result, err = client.Graph().CreateWithContext(context.Background(), gci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 値をピクセルに記録する
	pci := &pixela.PixelCreateInput{
		Date:     pixela.String("20180915"),
		Quantity: pixela.String("5"),
		GraphID:  pixela.String("graph-id"),
	}
	result, err = client.Pixel().CreateWithContext(context.Background(), pci)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}

	// 新しい webhook を作る
	wci := &pixela.WebhookCreateInput{
		GraphID: pixela.String("graph-id"),
		Type:    pixela.String(pixela.WebhookTypeIncrement),
	}
	webhook, err := client.Webhook().CreateWithContext(context.Background(), wci)
	if err != nil {
		log.Fatal(err)
	}
	if webhook.IsSuccess == false {
		log.Fatal(webhook.Message)
	}

	// webhook を呼び出す
	wii := &pixela.WebhookInvokeInput{WebhookHash: pixela.String("webhook-hash")}
	result, err = client.Webhook().InvokeWithContext(context.Background(), wii)
	if err != nil {
		log.Fatal(err)
	}
	if result.IsSuccess == false {
		log.Fatal(result.Message)
	}
}
```

## コントリビューション

1. このリポジトリをフォークします
2. issue ブランチを作成します (`git checkout -b issue/:id`)
3. コードを変更します
4. `make test` でテストを実行し, パスすることを確認します
5. `make fmt` でコードをフォーマットします
6. [Conventional Commits](https://www.conventionalcommits.org/) の形式で変更をコミットします
7. 新しいプルリクエストを作成します

### コミットの種類

| Type | 説明 |
|---|---|
| `feat` | 新しい API メソッドや入力フィールドの追加 |
| `fix` | 誤った動作の修正（URL・HTTP メソッド・フィールド処理の誤りなど） |
| `docs` | ドキュメントのみの変更（README・AGENTS.md・godoc コメント） |
| `refactor` | 動作を変えない内部構造の整理 |
| `test` | プロダクションコードを変更しないテストの追加・修正 |
| `chore` | 依存ライブラリの更新・置き換え、標準ライブラリへの移行、Makefile・go.mod の変更 |
| `ci` | CI 設定の変更（`.github/workflows/`） |

破壊的変更の場合は、type の後に `!` を付けるか（`feat!:`）、フッターに `BREAKING CHANGE:` を記載してください。

## ライセンス

[MIT](https://github.com/ebc-2in2crc/pixela4go/blob/master/LICENSE)

## 作者

[ebc-2in2crc](https://github.com/ebc-2in2crc)
