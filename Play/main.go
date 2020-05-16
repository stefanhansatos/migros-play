package main

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

	var unknownData map[string]interface{}
	err := json.Unmarshal(message.Data, &unknownData)

	fmt.Printf("unknownData: %v\n%t\n", unknownData["translatedText"], unknownData["translatedText"])
	if err != nil {
		return fmt.Errorf("failed to unmarshal unknownData: %v", err)
	}

	var translation *Translation
	err = json.Unmarshal(message.Data, &translation)

	var translationJson []byte
	translationJson, err = json.Marshal(translation)

	fmt.Printf("translationJson: %s\n", translationJson)

	wrappedTranslation := WrappedTranslation{
		Source: "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationLoad-europe-west1-smbe_output",
		LogFilter: fmt.Sprintf("resource.type=%q resource.labels.function_name=%q resource.labels.region=%q, europe-west1 severity=DEFAULT",
			"cloud_function", "SmbeTranslationLoad", "europe-west1"),
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

/*

  "jsonPayload": {
    "source": "projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslate-europe-west1-smbe_input",
    "unix": 1589629888,
    "logFilter": "xxx",
    "translation": {
      "translationQuery": {
        "sourceLanguage": "en",
        "targetLanguage": "fr",
        "text": "2: Tommorow is Tuesday"
      },
      "translatedText": "2: Tommorow est mardi",
      "translationErrors": [
        ""
      ]
    },
    "timestamp": "2020-05-16 11:51:28.966107881 +0000 UTC m=+0.922569518"
  }


{"source":"projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslate-europe-west1-smbe_input",
"translation":{"translationQuery":{"text":"4: Tommorow is Tuesday","sourceLanguage":"en","targetLanguage":"fr"},"translatedText":"4: Tommorow est mardi","translationErrors":[""]},
"logFilter":"xxx",
"timestamp":"2020-05-16 12:14:17.864054286 +0000 UTC m=+619.173984653",
"unix":1589631257}

*/
