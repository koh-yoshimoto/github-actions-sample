package main

import (
	"context"
	"encoding/json"
	"errors"
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
	if err != nil || req.Name == "" {
		err = errors.New("bad request")
		return Response{Message: "lambda function failed"}, err
	}

	return Response{Message: fmt.Sprintf("Hello %s from Lambda!", req.Name)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
