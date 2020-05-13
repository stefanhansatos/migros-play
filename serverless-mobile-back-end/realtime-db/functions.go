package realtime_db

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"fmt"
	"os"
	"time"
)

type Message struct {
	Data []byte `json:"data"`
}

type TranslationQuery struct {
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
	Timestamp   string       `json:"timestamp,omitempty"`
	Unix        int64        `json:"unix,omitempty"` // Unix time in seconds
}

type WrappedTranslationQuery struct {
	Source           string            `json:"source"`
	TranslationQuery *TranslationQuery `json:"translationQuery"`
	Timestamp        string            `json:"timestamp,omitempty"`
	Unix             int64             `json:"unix,omitempty"` // Unix time in seconds
}

// SmbeTranslationQueryLoad stores the translation query in smbe:translation-queries to store the pubsub message
func SmbeTranslationQueryLoad(ctx context.Context, message Message) error {

	var translationQuery *TranslationQuery
	err := json.Unmarshal(message.Data, &translationQuery)

	wrappedTranslationQuery := WrappedTranslationQuery{
		Source:           "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationQueryLoad-europe-west1-smbe_input",
		TranslationQuery: translationQuery,
		Timestamp:        time.Now().String(),
		Unix:             time.Now().Unix(),
	}

	databaseURL := os.Getenv("RTDB_URL")
	if databaseURL == "" {
		return fmt.Errorf("RTDB_URL not set")
	}

	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)

	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database client: %v", err)

	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/translation/queries")
	_, err = ref.Push(ctx, interface{}(&wrappedTranslationQuery))
	if err != nil {
		return fmt.Errorf("error pushing new list node: %v", err)

	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)
	return nil
}

// SmbeTranslationQueryLoad stores the translation query in smbe:translation-queries to store the pubsub message
func SmbeTranslationLoad(ctx context.Context, message Message) error {

	var translation *Translation
	err := json.Unmarshal(message.Data, &translation)

	wrappedTranslation := WrappedTranslation{
		Source:      "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationLoad-europe-west1-smbe_output",
		Translation: translation,
		Timestamp:   time.Now().String(),
		Unix:        time.Now().Unix(),
	}

	databaseURL := os.Getenv("RTDB_URL")
	if databaseURL == "" {
		return fmt.Errorf("RTDB_URL not set")
	}

	conf := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)

	}

	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database client: %v", err)

	}

	// As an admin, the app has access to read and write all data, regradless of Security Rules
	ref := client.NewRef("/translation/results")
	_, err = ref.Push(ctx, interface{}(&wrappedTranslation))
	if err != nil {
		return fmt.Errorf("error pushing new list node: %v", err)

	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)
	return nil
}
