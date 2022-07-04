# go-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/limoo-im/go-sdk)](https://goreportcard.com/report/github.com/limoo-im/go-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/limoo-im/go-sdk.svg)](https://pkg.go.dev/github.com/limoo-im/go-sdk)

An SDK for Limoo Written in Go

## Example

With following sample code you can use this SDK and send a message to Limoo

```go
package main

import (
	"github.com/limoo-im/go-sdk"
	"github.com/limoo-im/go-sdk/types"
)

func main() {
	client := &sdk.LimooClient{}
	err := client.New("https://web.limoo.im", "<username>", "<password>", false)
	if err != nil {
		panic(err)
	}
	sdk.SetDebug(true)
	err = client.SendMessage(types.SendMessageOptions{
		Text:           "Hello World",
		WorkspaceID:    "<ID of the workspace>",
		ConversationID: "<ID of the conversation>",
	})
	if err != nil {
		panic(err)
	}
}
```
