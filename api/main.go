package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type Payload struct {
	Body string `json:"body"`
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, payload Payload) (Response, error) {
	var req Request
	err := json.NewDecoder(strings.NewReader(payload.Body)).Decode(&req)
	if err != nil {
		return Response{Message: "lambda function failed"}, err
	}

	return Response{Message: fmt.Sprintf("Hello %s from Lambda!", req.Name)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
