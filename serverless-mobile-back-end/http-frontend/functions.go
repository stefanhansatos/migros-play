package http_frontend

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
)

type TranslationQuery struct {
	Uuid           string `json:"uuid"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

// SmbeHTTP is an entry point for the smbe
func SmbeHTTP(w http.ResponseWriter, r *http.Request) {

	firebaseProject := os.Getenv("FIREBASE_PROJECT")
	if firebaseProject == "" {
		http.Error(w, fmt.Sprintf("FIREBASE_PROJECT not set\n"), http.StatusInternalServerError)
		return

	}

	pubsubTopic := os.Getenv("SMBE_PUBSUB_TOPIC_IN")
	if pubsubTopic == "" {
		http.Error(w, fmt.Sprintf("SMBE_PUBSUB_TOPIC_IN not set\n"), http.StatusInternalServerError)
		return
	}

	var translationQuery TranslationQuery
	if err := json.NewDecoder(r.Body).Decode(&translationQuery); err != nil {
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to decode translationQuery: %v\n", err), http.StatusInternalServerError)
			return
		}
	}

	id, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create new random UUID: %v\n", err), http.StatusInternalServerError)
		return
	}
	translationQuery.Uuid = id.String()

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, firebaseProject)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create new pubsub client: %v\n", err), http.StatusInternalServerError)
		return
	}

	var translationJson []byte
	translationJson, err = json.Marshal(translationQuery)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal translationQuery: %v\n", err), http.StatusInternalServerError)
		return
	}

	topic := pubsubClient.Topic(pubsubTopic)
	defer topic.Stop()
	var results []*pubsub.PublishResult
	res := topic.Publish(ctx, &pubsub.Message{
		Data: translationJson,
	})
	results = append(results, res)
	// Do other work ...
	for _, r := range results {
		messageId, err := r.Get(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get pubsub result: %v\n", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Published a message with message ID %q and an internal UUID %q\n", messageId, id)
	}
}
