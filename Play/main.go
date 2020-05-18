package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
)

type TranslationQuery struct {
	Uuid           string `json:"uuid"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

type Translation struct {
	TranslationQuery  *TranslationQuery `json:"translationQuery"`
	TranslatedText    string            `json:"translatedText"`
	TranslationErrors []string          `json:"translationErrors"`
}

type WrappedTranslation struct {
	Source      string       `json:"source"`
	Translation *Translation `json:"translation"`
	LogFilter   string       `json:"logFilter"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Unix        int64        `json:"unix,omitempty"` // Unix time in seconds
}

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

	translationQuery := TranslationQuery{
		Uuid:           "uuid",
		Text:           "translate me",
		SourceLanguage: "en",
		TargetLanguage: "de",
	}
	translation := Translation{
		TranslationQuery:  &translationQuery,
		TranslatedText:    "uebersetz mich",
		TranslationErrors: []string{""},
	}
	wrappedTranslation := WrappedTranslation{
		Source:      "play",
		Translation: &translation,
		Timestamp:   time.Now().String(),
		Unix:        int64(1),
	}

	translations := []WrappedTranslation{
		wrappedTranslation,
	}

	smbeDataset := client.Dataset("smbe")
	translationsTable := smbeDataset.Table("translations")

	inserter := translationsTable.Inserter()
	err = inserter.Put(ctx, translations)
	if err != nil {
		log.Fatalf("inserter.Put: %v", err)
	}
}
