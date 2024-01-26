package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	ec2Client := ec2.New(ec2.Options{Credentials: cfg.Credentials, Region: "sa-east-1"})
	instances, err := ec2Client.DescribeInstances(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("InstanceID: %s, PublicIPAddress: %s\n", *instance.InstanceId, *instance.PublicIpAddress)
		}
	}
}
