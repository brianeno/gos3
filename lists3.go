package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// List: /list?prefix={prefix}&delimeter={delimeter}
func handlerList(w http.ResponseWriter, r *http.Request) {

	// There aren't really any folders in S3, but we can emulate them by using "/" in the key names
	// of the objects. In case we want to listen the contents of a "folder" in S3, what we really need
	// to do is to list all objects which have a certain prefix.
	prefix := r.URL.Query().Get("prefix")
	delimeter := r.URL.Query().Get("delimeter")

	paginator := s3.NewListObjectsV2Paginator(awsS3Client, &s3.ListObjectsV2Input{
		Bucket:    aws.String(AWS_S3_BUCKET),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String(delimeter),
	})

	w.Header().Set("Content-Type", "text/html")

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			fmt.Println("Error during list")
		}
		for _, obj := range page.Contents {
			// return come html
			fmt.Fprintf(w, "<li>File %s</li>", *obj.Key)
		}
	}
}
