package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// /create - /creates3 - create the bucket
func handlerCreateBucket(w http.ResponseWriter, r *http.Request) {

	_, err := awsS3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(AWS_S3_BUCKET),
	})

	if err != nil {
		showError(w, r, http.StatusInternalServerError, "Error during creating the bucket "+AWS_S3_BUCKET+" "+err.Error())
		return
	}

	fmt.Fprintf(w, "Successfully created bucket %s\n", AWS_S3_BUCKET)
}
