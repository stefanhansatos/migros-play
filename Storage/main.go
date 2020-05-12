package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {

	gcpProject := os.Getenv("GCP_PROJECT")
	if gcpProject == "" {
		fmt.Printf("GCP_PROJECT not set\n")
		return
	}

	storageCredentialFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if storageCredentialFile == "" {
		fmt.Printf("STORAGE_APPLICATION_CREDENTIALS not set\n")
		return
	}

	bucketUrl := os.Getenv("FIREBASE_BUCKET_URL")
	if bucketUrl == "" {
		fmt.Printf("FIREBASE_BUCKET_URL not set\n")
		return
	}
	fmt.Printf("FIREBASE_BUCKET_URL: %q\n", bucketUrl)
	config := &firebase.Config{
		StorageBucket: bucketUrl,
	}
	opt := option.WithCredentialsFile(storageCredentialFile)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("failed to create new firebase app: %v\n", err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf("failed to return storage instance: %v\n", err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalf("failed to return default bucket handle: %v\n", err)
	}

	ctx := context.Background()
	f, err := os.Open("../data/notes.txt")
	if err != nil {
		log.Fatalf("os.Open: %v\n", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := bucket.Object("notes.txt").NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Fatalf("io.Copy(wc, f): %v\n", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("wc.Close(): %v\n", err)
	}

	bucket, err = client.Bucket("my-custom-bucket-1234")
	if err != nil {
		log.Fatalf("failed to return custom bucket handle: %v\n", err)
	}
	err = bucket.Create(ctx, "hybrid-cloud-22365", nil)
	if err != nil {
		log.Fatalf("failed to create custom bucket: %v\n", err)
	}
	wc = bucket.Object("notes.txt").NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Fatalf("io.Copy(wc, f): %v\n", err)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("wc.Close(): %v\n", err)
	}
}
