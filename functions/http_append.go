package functions

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

// Append creates a new node in someData/list
func AppendHttp(w http.ResponseWriter, r *http.Request) {

	// Get rid of warnings
	_ = r

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

	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://hybrid-cloud-22365.firebaseio.com",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("./hybrid-cloud-22365-firebase-adminsdk-ca37q-d1e808e47b.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app: ", err)
		//fmt.Errorf("Error initializing app: %v", err)
		//return
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client: ", err)
		//fmt.Errorf("Error initializing database client: %v", err)
		//return
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
	newRef, err := ref.Push(ctx, interface{}(&someData))
	if err != nil {
		log.Fatalln("Error pushing new list node: ", err)
		//fmt.Errorf("Error pushing new list node: %v", err)
		//return
	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)

	fmt.Fprintf(w, "pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)

}
