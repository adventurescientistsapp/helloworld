package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type BasicEvent struct {
        Input string `json:"input"`
}

func HandleRequest(ctx context.Context, event BasicEvent) (string, error) {
        return fmt.Sprintf("Hello World! event.Input=%s", event.Input ), nil
}

func main() {
        lambda.Start(HandleRequest)
}