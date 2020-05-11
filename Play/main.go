package main

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

type Payload struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type wrappedData struct {
	Source    string   `json:"source"`
	Payload   *Payload `json:"payload"`
	Name      string   `json:"name,omitempty"`
	Number    int      `json:"number,omitempty"`
	Desc      string   `json:"description,omitempty"`
	Status    string   `json:"status,omitempty"`
	Timestamp string   `json:"timestamp,omitempty"`
	Unix      int64    `json:"unix,omitempty"` // Unix time in seconds
}

func main() {

	payload := Payload{
		Type: "string",
		Data: 11,
	}

	payloadBytes, err := json.Marshal(payload)
	fmt.Printf("payloadBytes: %v\n", string(payloadBytes))

	var createdPayload *Payload
	err = json.Unmarshal(payloadBytes, &createdPayload)
	fmt.Printf("createdPayload: %v\n", createdPayload)

	someData := wrappedData{
		Source:    "/Users/stefan/go/src/github.com/stefanhansatos/migros-play/Play/main.go",
		Payload:   createdPayload,
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

}
