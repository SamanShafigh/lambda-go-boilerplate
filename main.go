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

// Declaring lambda app
var appInstanse *app.App

// HandleRequest handles events
func HandleRequest(ctx context.Context, event MyEvent) (string, error) {

	res, err := appInstanse.Run()
	if err != nil {
		return fmt.Sprintf("There are some errors during run time [%s]", err), nil
	}

	return fmt.Sprintf("Hello %s! this is APP res: [%s]", event.Name, res), nil
}

// HandleError handles errors in app
func HandleError(message string, err error) string {
	return fmt.Sprintf("%s [%s]", message, err)
}

// The entry point of app
func main() {
	var err error
	// Initialising a new version of the application
	appInstanse, err = app.New(ConfigPath)
	if err != nil {
		panic(err)
	}

	// Close the database connection when main is finish
	defer appInstanse.Model.DB.Close()

	// Start Lambda with this handler
	lambda.Start(HandleRequest)
}
