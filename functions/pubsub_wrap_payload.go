package functions

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
	"os"
	"time"
)

// Store creates a new node in someData/list to store the pubsub message
func WrapPayload(ctx context.Context, message Message) error {

	var payload *Payload
	err := json.Unmarshal(message.Data, &payload)

	wrappedData := WrappedData{
		Source:    "projects/hybrid-cloud-22365/subscriptions/gcf-WrapPayload-europe-west1-fb_someData",
		Payload:   payload,
		Timestamp: time.Now().String(),
		Unix:      time.Now().Unix(),
	}

	databaseURL := os.Getenv("FIREBASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("FIREBASE_URL not set")
	}

	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)

	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database client: %v", err)

	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/someData/list")
	_, err = ref.Push(ctx, interface{}(&wrappedData))
	if err != nil {
		return fmt.Errorf("error pushing new list node: %v", err)

	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)
	return nil
}
