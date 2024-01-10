package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	AWS_S3_REGION = "us-east-1"          // Region
	AWS_S3_BUCKET = "brianeno-go-bucket" // Bucket
)

func main() {
	configS3()

	http.HandleFunc("/uploads3", handlerUpload)
	http.HandleFunc("/downloads3", handlerDownload)
	http.HandleFunc("/lists3", handlerList)
	http.HandleFunc("/creates3", handlerCreateBucket)
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func showError(w http.ResponseWriter, r *http.Request, status int, message string) {
	http.Error(w, message, status)
}
