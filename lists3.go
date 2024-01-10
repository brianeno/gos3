package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// List: /lists3?prefix={prefix}&delimeter={delimeter}
func handlerList(w http.ResponseWriter, r *http.Request) {

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
		} else {
			for _, obj := range page.Contents {
				// return come html
				fmt.Fprintf(w, "<li>File %s</li>", *obj.Key)
			}
		}
	}
}
