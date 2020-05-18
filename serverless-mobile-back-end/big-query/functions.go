package big_query

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"os"
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

type WrappedTranslationQuery struct {
	Source           string            `json:"source"`
	TranslationQuery *TranslationQuery `json:"translationQuery"`
	Timestamp        string            `json:"timestamp,omitempty"`
	Unix             int64             `json:"unix,omitempty"` // Unix time in seconds
}

type WrappedData struct {
	Source      string       `json:"source"`
	Translation *Translation `json:"translation"`
	LogFilter   string       `json:"logFilter"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Unix        int64        `json:"unix,omitempty"` // Unix time in seconds
}

// SmbeBqLoad stores the translation in BigQuery "smbe:translations"
func SmbeBqLoad(ctx context.Context, message pubsub.Message) error {

	// resource.type="cloud_function" resource.labels.function_name="SmbeTranslationLoad" resource.labels.region="europe-west1" severity=DEFAULT

	var wrappedData WrappedData
	err := json.Unmarshal(message.Data, &wrappedData)

	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		return fmt.Errorf("GCP_PROJECT not set")
	}

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create bigquery client: %v", err)
	}
	defer client.Close()

	translations := []WrappedData{
		wrappedData,
	}

	smbeDataset := client.Dataset("smbe")
	translationsTable := smbeDataset.Table("translations")

	inserter := translationsTable.Inserter()
	err = inserter.Put(ctx, translations)
	if err != nil {
		return fmt.Errorf("failed to insert data into bigquery: %v", err)
	}
	return nil
}
