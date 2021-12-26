package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	var instanceId string
	var action string

	flag.StringVar(&instanceId, "id", "", "Instance id")
	flag.StringVar(&action, "action", "start", "Action to perform [start/stop]")
	flag.Parse()

	if instanceId == "" {
		fmt.Println("No instance id provided.")
		os.Exit(1)
	}

	sess := session.Must(session.NewSession())

	svc := ec2.New(sess)

	switch action {
	case "start":
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(instanceId),
			},
		}

		response, err := svc.StartInstances(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				fmt.Println(aerr.Error())
			} else {
				fmt.Println(err.Error())
			}

			return
		}

		fmt.Println(response)

	case "stop":
		input := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(instanceId),
			},
		}

		response, err := svc.StopInstances(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				fmt.Println(aerr.Error())
			} else {
				fmt.Println(err.Error())
			}

			return
		}

		fmt.Println(response)

	default:
		fmt.Printf("Invalid action '%s'\n", action)
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

}
