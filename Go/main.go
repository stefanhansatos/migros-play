package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"

	"firebase.google.com/go"
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

func main() {

	someData := SomeData{
		Name:      "Alice",
		Number:    42,
		Desc:      "Only test data to play with",
		Status:    "None",
		Timestamp: time.Now().String(),
		Unix:      time.Now().Unix(),
	}

	fmt.Printf("%v\n", someData)

	// Marshall someData
	jsonData, err := json.Marshal(someData)
	if err != nil {
		fmt.Printf("failed to unmarshall %q: %v\n", someData, err)
		return
	}
	fmt.Printf("someData: %s\n", jsonData)

	firebaseCredentialFile := os.Getenv("FIREBASE_APPLICATION_CREDENTIALS")
	if firebaseCredentialFile == "" {
		fmt.Printf("FIREBASE_APPLICATION_CREDENTIALS not set\n")
		return
	}

	firebaseProject := os.Getenv("FIREBASE_PROJECT")
	if firebaseProject == "" {
		fmt.Printf("FIREBASE_PROJECT not set\n")
		return
	}

	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: fmt.Sprintf("https://%s.firebaseio.com", firebaseProject),
		ProjectID:   firebaseProject,
		// ServiceAccountID: "113262279650432319262",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(firebaseCredentialFile)

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/someData/list")
	newRef, err := ref.Push(ctx, interface{}(&someData))
	if err != nil {
		log.Fatalf("Error pushing new list node: %v", err)
	}
	fmt.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)

	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	//fmt.Println(data)

	jsonData, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("failed to unmarshall %q: %v\n", data, err)
		return
	}
	fmt.Printf("someData: %s\n", jsonData)

}
