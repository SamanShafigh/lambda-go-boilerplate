package main

import (
	"context"
	"fmt"

	"github.com/SamanShafigh/lambda-go-boilerplate/app"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	// ConfigPath defines the path to app config file
	ConfigPath = "conf.json"
)

// MyEvent is incomming event structure
type MyEvent struct {
	Name string `json:"name"`
}

// HandleRequest handles events
func HandleRequest(ctx context.Context, event MyEvent) (string, error) {

	app, err := app.New(ConfigPath)
	if err != nil {
		return HandleError("There are some errors during app initialization", err), nil
	}

	res, err := app.Run()
	if err != nil {
		return HandleError("There are some errors during run time", err), nil
	}

	return fmt.Sprintf("Hello %s! this is APP res: [%s]", event.Name, res), nil
}

// HandleError handles errors in app
func HandleError(message string, err error) string {
	return fmt.Sprintf("%s [%s]", message, err)
}

// The entry point of app
func main() {
	lambda.Start(HandleRequest)
}
