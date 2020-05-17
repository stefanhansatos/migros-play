package storage

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
	"time"

	firebase "firebase.google.com/go"
)

func SmbeFileStore(ctx context.Context, message pubsub.Message) error {

	firebaseProject := os.Getenv("FIREBASE_PROJECT")
	if firebaseProject == "" {
		return fmt.Errorf("FIREBASE_PROJECT not set\n")

	}

	bucketUrl := os.Getenv("FIREBASE_BUCKET_URL")
	if bucketUrl == "" {
		return fmt.Errorf("FIREBASE_BUCKET_URL not set\n")
	}
	fmt.Printf("FIREBASE_BUCKET_URL: %q\n", bucketUrl)
	config := &firebase.Config{
		StorageBucket: bucketUrl,
	}
	//opt := option.WithCredentialsFile(storageCredentialFile)
	//app, err := firebase.NewApp(context.Background(), config, opt)
	//if err != nil {
	//	log.Fatalf("failed to create new firebase app: %v\n", err)
	//}

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		return fmt.Errorf("failed to create new firebase app: %v\n", err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return fmt.Errorf("failed to return storage instance: %v\n", err)
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		return fmt.Errorf("failed to return default bucket handle: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	var fileName = message.Attributes["uuid"]
	fmt.Printf("message.ID: %q\n", message.ID)
	fmt.Printf("message.Attributes: %v\n", message.Attributes)

	if fileName == "" {
		fileName = "directcall"
	}

	wc := bucket.Object(fileName).NewWriter(ctx)
	_, err = wc.Write(message.Data)
	if err != nil {
		return fmt.Errorf("failed to write to bucket: %v\n", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to close bucket handle: %v\n", err)
	}

	return nil
}
