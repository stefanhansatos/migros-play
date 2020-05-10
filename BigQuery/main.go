package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func main() {
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		fmt.Println("GCP_PROJECT environment variable must be set.")
		os.Exit(1)
	}

	bigQueryCredentialFile := os.Getenv("BIGQUERY_APPLICATION_CREDENTIALS")
	if bigQueryCredentialFile == "" {
		fmt.Printf("BIGQUERY_APPLICATION_CREDENTIALS not set\n")
		return
	}

	ctx := context.Background()

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(bigQueryCredentialFile)

	client, err := bigquery.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	rows, err := query(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
	if err := printResults(os.Stdout, rows); err != nil {
		log.Fatal(err)
	}
}

// query returns a row iterator suitable for reading query results.
func query(ctx context.Context, client *bigquery.Client) (*bigquery.RowIterator, error) {

	query := client.Query(
		`SELECT name,count FROM babynames.names2010 WHERE gender = 'F' ORDER BY count DESC LIMIT 5`)
	return query.Read(ctx)
}

// [START bigquery_simple_app_print]
type ResultRow struct {
	Name  string `bigquery:"name"`
	Count int64  `bigquery:"count"`
}

// printResults prints results from a query to the Stack Overflow public dataset.
func printResults(w io.Writer, iter *bigquery.RowIterator) error {
	for {
		var row ResultRow
		err := iter.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error iterating through results: %v", err)
		}

		fmt.Fprintf(w, "Name: %s\t Count: %d\n", row.Name, row.Count)
	}
}
