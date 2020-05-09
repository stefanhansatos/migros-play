package functions

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

type SomeData struct {
	ID        string `json:"id"` // ID from Firebase DB
	Name      string `json:"name,omitempty"`
	Number    int    `json:"number,omitempty"`
	Desc      string `json:"description,omitempty"`
	Status    string `json:"status,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Unix      int64  `json:"unix,omitempty"` // Unix time in seconds
}

// Append creates a new node in someData/list
func Append(ctx context.Context, r *http.Request) error {

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

	//ctx := context.Background()
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

	someData := SomeData{
		Name:      "Alice",
		Number:    42,
		Desc:      "Only test data to play with",
		Status:    "None",
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

// gcloud pubsub topics create fb_someData
// gcloud functions deploy Append --region europe-west1 --runtime go111 --trigger-topic=fb_someData
// gcloud functions call Append --region europe-west1 --data '{}'

// gcloud pubsub topics publish fb_someData --message "not used"

/*
{
"rules": {
"users": {
"$uid": {
".read": "$uid === auth.uid",
".write": "$uid === auth.uid"
}
}
}
}
*/
