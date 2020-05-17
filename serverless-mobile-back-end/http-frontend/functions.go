package http_frontend

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TranslationQuery struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func SmbeHTTP(w http.ResponseWriter, r *http.Request) {

	var translationQuery TranslationQuery
	if err := json.NewDecoder(r.Body).Decode(&translationQuery); err != nil {
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to decode translationQuery: %v\n", err), http.StatusInternalServerError)
			return
		}

	}

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "hybrid-cloud-22365")
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

	topic := pubsubClient.Topic("smbe_input")
	defer topic.Stop()
	var results []*pubsub.PublishResult
	res := topic.Publish(ctx, &pubsub.Message{
		Data: translationJson,
	})
	results = append(results, res)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get pubsub result: %v\n", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Published a message with a message ID: %s\n", id)
	}

}
