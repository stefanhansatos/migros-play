package functions

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func BqQuery(ctx context.Context, m interface{}) error {

	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		fmt.Println("GCP_PROJECT environment variable must be set.")
		os.Exit(1)
	}

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("./hybrid-cloud-22365-firebase-bq-22365.json")

	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	rows, err := query(ctx, client)
	if err != nil {
		return fmt.Errorf("bigquery.query: %v", err)
	}
	for {
		var row ResultRow
		err := rows.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("error iterating through results: %v", err)
		}
		log.Printf("Name: %s\t Count: %d\n", row.Name, row.Count)
	}

	return nil
}
