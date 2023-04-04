package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: "Hello, " + request.Name,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
