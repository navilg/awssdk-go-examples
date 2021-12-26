package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	var tagKey string
	var tagVal string

	flag.StringVar(&tagKey, "tag", "", "Tag Key")
	flag.StringVar(&tagVal, "value", "", "Tag Value")

	flag.Parse()

	sess := session.Must(session.NewSession())

	svc := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + tagKey),
				Values: []*string{
					aws.String(tagVal),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		return
	}

	// fmt.Println(result.Reservations)

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println(*instance.InstanceId)
		}
	}

}
