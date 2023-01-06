package main

import (
	"context"
	"fmt"

	"dupyeeter/internal/config"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		// return 500
		panic("catch errors better")
	}
	fmt.Printf("cfg %s", cfg.Strava.ClientID)
	lambda.Start(HandleRequest)
}
