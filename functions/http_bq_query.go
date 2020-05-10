package functions

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func Http_Query(w http.ResponseWriter, r *http.Request) {

	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		fmt.Println("GCP_PROJECT environment variable must be set.")
		os.Exit(1)
	}

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("./hybrid-cloud-22365-firebase-bq-22365.json")

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	rows, err := query(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var row ResultRow
		err := rows.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(fmt.Errorf("error iterating through results: %v", err))
		}

		fmt.Fprintf(w, "Name: %s\t Count: %d\n", row.Name, row.Count)
	}
}

// query returns a row iterator suitable for reading query results.
func query(ctx context.Context, client *bigquery.Client) (*bigquery.RowIterator, error) {

	query := client.Query(
		`SELECT name,count FROM babynames.names2010 WHERE gender = 'F' ORDER BY count DESC LIMIT 5`)
	return query.Read(ctx)
}

type ResultRow struct {
	Name  string `bigquery:"name"`
	Count int64  `bigquery:"count"`
}
