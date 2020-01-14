package main

import (
	"encoding/json"
	"fmt"
	"github.com/AntoineViau/lbcfb/fizzbuzz"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := request.QueryStringParameters
	for _, paramName := range []string{"string1", "string2", "int1", "int2", "limit"} {
		if params[paramName] == "" {
			return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Missing parameter %s", paramName), StatusCode: 400}, nil
		}
	}
	string1 := params["string1"]
	string2 := params["string2"]
	int1, err := strconv.Atoi(params["int1"])
	if err != nil {
		return error400("Invalid parameter int1"), nil
	}
	int2, err := strconv.Atoi(params["int2"])
	if err != nil {
		return error400("Invalid parameter int2"), nil
	}
	limit, err := strconv.Atoi(params["limit"])
	if err != nil || limit < 0 {
		return error400("Invalid parameter limit"), nil
	}

	output, err := fizzbuzz.FizzBuzz(string1, string2, int1, int2, limit)
	json, err := json.Marshal(output)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Internal FizzBuzz error", StatusCode: 500}, nil
	}
	headers := map[string]string{"Access-Control-Allow-Origin": "*"}
	return events.APIGatewayProxyResponse{Body: string(json), Headers: headers, StatusCode: 200}, nil
}

func error400(message string) events.APIGatewayProxyResponse {
	headers := map[string]string{"Access-Control-Allow-Origin": "*"}
	return events.APIGatewayProxyResponse{Body: message, Headers: headers, StatusCode: 400}
}

func main() {
	lambda.Start(Handler)
}
