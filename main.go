package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func handler(request events.APIGatewayProxyRequest) {
	fmt.Println(request)
}

func main() {
	
}
