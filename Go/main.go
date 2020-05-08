package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"

	"firebase.google.com/go"
)

func main() {

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
	ref := client.NewRef("/")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)

	ref.Push()

}
