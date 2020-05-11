// Triggers Pub/Sub topic with some data
package main

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"os"
)

type SomeData struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	Number    int    `json:"number,omitempty"`
	Desc      string `json:"description,omitempty"`
	Status    string `json:"status,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Unix      int64  `json:"unix,omitempty"` // Unix time in seconds
}

func main() {

	//env := os.Environ()

	//someData := SomeData{
	//	Name:      "From main.go: Triggers Pub/Sub topic with some data",
	//	Number:    42,
	//	Desc:      "Only test data to play with",
	//	Status:    fmt.Sprintf("{ %q: %q }", "env", "value"),
	//	Timestamp: time.Now().String(),
	//	Unix:      time.Now().Unix(),
	//}
	//
	//fmt.Printf("%v\n", someData)
	//
	//// Marshall someData
	//jsonData, err := json.Marshal(someData)
	//if err != nil {
	//	fmt.Printf("failed to unmarshall %q: %v\n", someData, err)
	//	return
	//}
	//fmt.Printf("someData: %s\n", jsonData)

	firebaseCredentialFile := os.Getenv("FIREBASE_APPLICATION_CREDENTIALS")
	if firebaseCredentialFile == "" {
		fmt.Printf("FIREBASE_APPLICATION_CREDENTIALS not set\n")
		return
	}

	gcpProject := os.Getenv("GCP_PROJECT")
	if gcpProject == "" {
		fmt.Printf("GCP_PROJECT not set\n")
		return
	}

	ctx := context.Background()

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("/Users/stefan/.secret/hybrid-cloud-22365-firebase-pubsub-22365.json")
	client, err := pubsub.NewClient(ctx, gcpProject, opt)
	if err != nil {
		fmt.Printf("failed to create pubsub client for project %q: %v\n", gcpProject, err)
		return
	}

	topic := client.Topic("fb_someData")
	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(fmt.Sprintf("{ %q: 1 }", "payload")),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			fmt.Printf("failed to create pubsub client: %v\n", err)
			return
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
}

// gcloud projects add-iam-policy-binding ${GCP_PROJECT}  --member serviceAccount:firebase-pubsub-22365@hybrid-cloud-22365.iam.gserviceaccount.com --role roles/pubsub.publisher
