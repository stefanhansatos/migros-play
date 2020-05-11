package main

import (
	//"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// ==================================================================
// https://firebase.google.com/docs/storage/admin/start
// ==================================================================

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

	//bucket, err := client.DefaultBucket()
	//if err != nil {
	//	log.Fatalf("client.DefaultBucket(): %v\n", err)
	//}
	//// 'bucket' is an object defined in the cloud.google.com/go/Storage package.
	//// See https://godoc.org/cloud.google.com/go/storage#BucketHandle
	//// for more details.
	//
	//log.Printf("Created bucket handle: %v\n", bucket)

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

	bucket, err = client.Bucket("my-custom-bucket-1324")
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

	return

	//obj := bucket.Object("hybrid-cloud-test-rmme123")
	//
	//ctx := context.Background()
	//// Write something to obj.
	//// w implements io.Writer.
	//w := obj.NewWriter(ctx)
	//// Write some text to obj. This will either create the object or overwrite whatever is there already.
	//if _, err := fmt.Fprintf(w, "This object contains text.\n"); err != nil {
	//	log.Fatalf("fmt.Fprintf: %v\n", err)
	//}
	//defer w.Close()
	//
	//
	//// Close, just like writing a file.
	//if err := w.Close(); err != nil {
	//	log.Fatalf("w.Close(): %v\n", err)
	//}
	//
	//
	//// Read it back.
	//r, err := obj.NewReader(ctx)
	//if err != nil {
	//	log.Fatalf("obj.NewReader(ctx): %v\n", err)
	//}
	//defer r.Close()
	//if _, err := io.Copy(os.Stdout, r); err != nil {
	//	log.Fatalf("io.Copy(os.Stdout, r): %v\n", err)
	//}
	//// Prints "This object contains text."
	//
	//return
	//if err := bucket.Create(ctx, gcpProject, nil); err != nil {
	//	fmt.Printf("failed to create bucket : err", err)
	//	return
	//}
	//
	//attrs, err := bucket.Attrs(ctx)
	//if err == storage.ErrBucketNotExist {
	//	fmt.Println("The bucket does not exist")
	//	return
	//}
	//if err != nil {
	//	fmt.Printf("Unknown error: err", err)
	//	return
	//}
	//fmt.Printf("The bucket exists and has attributes: %#v\n", attrs)
}
