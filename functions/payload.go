package functions

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"time"
)

type WrappedData struct {
	Payload   []byte      `json:"payload"`
	ID        string      `json:"id"`
	Name      string      `json:"name,omitempty"`
	Number    int         `json:"number,omitempty"`
	Desc      interface{} `json:"description,omitempty"`
	Status    string      `json:"status,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
	Unix      int64       `json:"unix,omitempty"` // Unix time in seconds
}

// Store creates a new node in someData/list to store the pubsub message
func Payload(ctx context.Context, m PubSubMessage) error {

	/*firebaseCredentialFile := os.Getenv("FIREBASE_APPLICATION_CREDENTIALS")
	if firebaseCredentialFile == "" {
		fmt.Printf("FIREBASE_APPLICATION_CREDENTIALS not set\n")
		return
	}

	firebaseProject := os.Getenv("FIREBASE_PROJECT")
	if firebaseProject == "" {
		fmt.Printf("FIREBASE_PROJECT not set\n")
		return
	}*/

	conf := &firebase.Config{
		DatabaseURL: "https://hybrid-cloud-22365.firebaseio.com",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("./hybrid-cloud-22365-firebase-adminsdk-ca37q-d1e808e47b.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return fmt.Errorf("Error initializing app: %v", err)

	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("Error initializing database client: %v", err)

	}
	name := string(m.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)

	wrappedData := WrappedData{
		Payload:   m.Data,
		Name:      string(m.Data),
		Number:    21,
		Desc:      interface{}(&m.Data),
		Status:    "None",
		Timestamp: time.Now().String(),
		Unix:      time.Now().Unix(),
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/someData/list")
	_, err = ref.Push(ctx, interface{}(&wrappedData))
	if err != nil {
		return fmt.Errorf("Error pushing new list node: %v", err)

	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)
	return nil
}
