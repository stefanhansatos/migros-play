package functions

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"log"
	"os"
	"time"
)

// Store creates a new node in someData/list to store the pubsub message
func Store(ctx context.Context, message Message) error {

	databaseURL := os.Getenv("FIREBASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("FIREBASE_URL not set")
	}

	env := os.Environ()
	envText := fmt.Sprint(env)

	//ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return fmt.Errorf("Error initializing app: %v", err)

	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("Error initializing database client: %v", err)

	}
	name := string(message.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)

	someData := SomeData{
		Name:      "projects/hybrid-cloud-22365/subscriptions/gcf-Store-europe-west1-fb_someData",
		Number:    21,
		Desc:      "pubsub_store.go receives data",
		Status:    envText,
		Timestamp: time.Now().String(),
		Unix:      time.Now().Unix(),
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/someData/list")
	_, err = ref.Push(ctx, interface{}(&someData))
	if err != nil {
		return fmt.Errorf("Error pushing new list node: %v", err)

	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)
	return nil
}
