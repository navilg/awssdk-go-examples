package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var bucket, key string
	var timeout time.Duration

	flag.StringVar(&bucket, "b", "", "Bucket Name")
	flag.StringVar(&key, "k", "", "Object key Name")
	flag.DurationVar(&timeout, "d", 0, "Upload timeout")
	flag.Parse()

	// Create a session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials.

	sess := session.Must(session.NewSession())

	// create service client for S3

	svc := s3.New(sess)

	// Create a context to setup timeout.

	ctx := context.Background() // Empty context with no timeout

	var cancelFn func()

	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}

	// CancelFn will run in background and if time exceeds timeout it stops the process defined in below code.
	defer cancelFn()

	_, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   os.Stdin,
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}

		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
}
