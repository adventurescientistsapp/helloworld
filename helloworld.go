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
  if event.Input == "" {
    return "Empty event.Input not allowed", nil
  }
  if event.Input == "version" {
    return "This is the second version of this lambda, uploaded via Travis CI + CD!!!", nil
  }	
  return fmt.Sprintf("Hello World! event.Input=%s", event.Input ), nil
}

func main() {
  lambda.Start(HandleRequest)
}
