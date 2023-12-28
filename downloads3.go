package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Download: /download?key={key of the object}&filename={name for the new file}
func handlerDownload(w http.ResponseWriter, r *http.Request) {

	// We get the name of the file on the URL
	filename := r.URL.Query().Get("filename")

	// We get the name of the file on the URL
	key := r.URL.Query().Get("key")

	// Create the file
	newFile, err := os.Create(filename)
	if err != nil {
		showError(w, r, http.StatusBadRequest, "Error creating the local file")
	}
	defer newFile.Close()

	downloader := manager.NewDownloader(awsS3Client)
	_, err = downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(key),
	})

	if err != nil {
		showError(w, r, http.StatusBadRequest, "Error retrieving the file from S3")
		return
	}

	http.ServeFile(w, r, filename)
}
