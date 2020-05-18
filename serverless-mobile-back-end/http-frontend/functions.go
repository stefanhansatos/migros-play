package http_frontend

import (
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/translate"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"net/http"
	"os"
)

type TranslationQuery struct {
	Uuid           string `json:"uuid"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

type Response struct {
	Uuid           string `json:"uuid"`
	TranslatedText string `json:"translatedText"`
	LoadCommand    string `json:"loadCommand"`
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
	translateClient, err := translate.NewClient(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create new translate client: %v\n", err), http.StatusInternalServerError)
		return
	}

	// Use the client.
	ds, err := translateClient.DetectLanguage(ctx, []string{translationQuery.Text})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to detect language: %v\n", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(ds)

	if ds[0][0].Language.String() != translationQuery.SourceLanguage {
		http.Error(w, fmt.Sprintf("source language is %q, but expected is %q\n", ds[0][0].Language.String(), translationQuery.SourceLanguage),
			http.StatusInternalServerError)
		return
	}

	if ds[0][0].Confidence < 0.9 {
		http.Error(w, fmt.Sprintf("source language detection's confidence for %q is below 90%\n", ds[0][0].Language.String()),
			http.StatusInternalServerError)
		return
	}

	langs, err := translateClient.SupportedLanguages(ctx, language.English)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to retrieve supported languages: %v\n", err), http.StatusInternalServerError)
		return
	}
	//fmt.Println(langs)

	var targetTag language.Tag
	for _, lang := range langs {
		if lang.Tag.String() == translationQuery.TargetLanguage {
			targetTag = lang.Tag
		}
	}

	translations, err := translateClient.Translate(ctx,
		[]string{translationQuery.Text}, targetTag,
		&translate.Options{
			Source: ds[0][0].Language,
			Format: translate.Text,
		})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to translate text: %v\n", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(translations[0].Text)

	response := Response{
		Uuid:           id.String(),
		TranslatedText: translations[0].Text,
		LoadCommand:    fmt.Sprintf("gsutil cat gs://hybrid-cloud-22365.appspot.com/%s | jq", id),
	}

	//ctx := context.Background()
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
		_ = messageId // future use?

		responseJson, err := json.MarshalIndent(response, "", " ")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal response: %v\n", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s\n", responseJson)

	}
}
