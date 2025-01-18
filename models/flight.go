package models

import (
	"arrivals-service/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Flight represents an arriving flight
type Flight struct {
	FlightID    string `json:"flight_id"`
	ArrivalTime string `json:"arrival_time"`
	Origin      string `json:"origin"`
	Status      string `json:"status"`
}

// FetchArrivingFlights retrieves flights from the DynamoDB table
func FetchArrivingFlights() ([]Flight, error) {
	svc := config.GetDynamoDBClient()
	input := &dynamodb.ScanInput{
		TableName: aws.String("ArrivingFlights"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		return nil, fmt.Errorf("error scanning table: %v", err)
	}

	var flights []Flight
	for _, item := range result.Items {
		flight := Flight{}
		err = dynamodbattribute.UnmarshalMap(item, &flight)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling item: %v", err)
		}
		flights = append(flights, flight)
	}

	return flights, nil
}
