package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Upload: /uploads3 (upload file named "file")
func handlerUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	// Get a file from the form input name "file"
	file, header, err := r.FormFile("file")
	if err != nil {
		showError(w, r, http.StatusInternalServerError, "Error during retrieving the file from the request form")
		return
	}
	defer file.Close()

	filename := header.Filename

	uploader := manager.NewUploader(awsS3Client)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		showError(w, r, http.StatusInternalServerError, "Error during uploading the file"+err.Error())
		return
	}

	fmt.Fprintf(w, "Successfully uploaded to %s\n", AWS_S3_BUCKET)
}
