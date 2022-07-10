# go-sdk

[![Go Report Card](https://goreportcard.com/badge/github.com/limoo-im/go-sdk)](https://goreportcard.com/report/github.com/limoo-im/go-sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/limoo-im/go-sdk.svg)](https://pkg.go.dev/github.com/limoo-im/go-sdk)

Limoo GoLang SDK

## Example

With following sample code you can use this SDK and send a message to Limoo:

```go
package main

import (
    "github.com/limoo-im/go-sdk"
    "github.com/limoo-im/go-sdk/types"
)

func main() {
    create empty client and inialize it
    client := &sdk.LimooClient{}
    err := client.New("https://web.limoo.im", "<username>", "<password>", false)
    if err != nil {
        panic(err)
    }
    
    // print debug log (will print sensitive data)
    sdk.SetDebug(true)
    
    // simply send a message
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

## License
This SDK licensed under [Apache-2 license](LICENSE).
