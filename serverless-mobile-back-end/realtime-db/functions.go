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

// SmbeTranslationQueryLoad stores the translation query in smbe:translation-queries to store the pubsub message
func SmbeTranslationQueryLoad(ctx context.Context, message Message) error {

	var translationQuery *TranslationQuery
	err := json.Unmarshal(message.Data, &translationQuery)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translationQuery: %v", err)
	}

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

	// resource.type="cloud_function" resource.labels.function_name="SmbeTranslationLoad" resource.labels.region="europe-west1" severity=DEFAULT

	fmt.Printf("message.Data: %s (%t)\n", message.Data, message.Data)
	var unknownData map[string]interface{}
	err := json.Unmarshal(message.Data, &unknownData)

	fmt.Printf("unknownData: %v\n%t\n", unknownData["translatedText"], unknownData["translatedText"])
	if err != nil {
		return fmt.Errorf("failed to unmarshal unknownData: %v", err)
	}

	var wrappedData *WrappedData
	err = json.Unmarshal(message.Data, &wrappedData)

	//var translation *Translation
	//err = json.Unmarshal(message.Data, &translation)
	//
	//var translationJson []byte
	//translationJson, err = json.Marshal(translation)
	//
	//fmt.Printf("translationJson: %s\n", translationJson)
	//
	//wrappedTranslation := WrappedTranslation{
	//	Source: "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationLoad-europe-west1-smbe_output",
	//	LogFilter: fmt.Sprintf("resource.type=%q resource.labels.function_name=%q resource.labels.region=%q, europe-west1 severity=DEFAULT",
	//		"cloud_function", "SmbeTranslationLoad", "europe-west1"),
	//	Translation: translation,
	//	Timestamp:   time.Now().String(),
	//	Unix:        time.Now().Unix(),
	//}
	//_ := wrappedTranslation

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
	_, err = ref.Push(ctx, interface{}(&wrappedData))
	if err != nil {
		return fmt.Errorf("error pushing new list node: %v", err)

	}
	//log.Printf("pushing new list node at %q: %v\n", newRef.Parent().Path, newRef.Key)
	return nil
}
