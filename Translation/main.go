package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

//func Example_NewClient() {
//	ctx := context.Background()
//	client, err := translate.NewClient(ctx)
//	if err != nil {
//		// TODO: handle error.
//	}
//	// Use the client.
//
//	// Close the client when finished.
//	if err := client.Close(); err != nil {
//		// TODO: handle error.
//	}
//}
//
//func Example_Translate() {
//	ctx := context.Background()
//	client, err := translate.NewClient(ctx)
//	if err != nil {
//		// TODO: handle error.
//	}
//	translations, err := client.Translate(ctx,
//		[]string{"Le singe est sur la branche"}, language.English,
//		&translate.Options{
//			Source: language.French,
//			Format: translate.Text,
//		})
//	if err != nil {
//		// TODO: handle error.
//	}
//	fmt.Println(translations[0].Text)
//}
//
//func Example_DetectLanguage() {
//	ctx := context.Background()
//	client, err := translate.NewClient(ctx)
//	if err != nil {
//		// TODO: handle error.
//	}
//	ds, err := client.DetectLanguage(ctx, []string{"Today is Monday"})
//	if err != nil {
//		// TODO: handle error.
//	}
//	fmt.Println(ds)
//}
//
//func Example_SupportedLanguages() {
//	ctx := context.Background()
//	client, err := translate.NewClient(ctx)
//	if err != nil {
//		// TODO: handle error.
//	}
//	langs, err := client.SupportedLanguages(ctx, language.English)
//	if err != nil {
//		// TODO: handle error.
//	}
//	fmt.Println(langs)
//}

func main() {

	text := "Today is Monday"
	sourceLanguage := language.English
	targetLanguage := language.French

	ctx := context.Background()
	translateClient, err := translate.NewClient(ctx)
	if err != nil {
		fmt.Printf("failed to create new translate client: %v\n", err)
		return
	}

	// Use the client.
	ds, err := translateClient.DetectLanguage(ctx, []string{text})
	if err != nil {
		fmt.Printf("failed to detect language: %v\n", err)
		return
	}
	fmt.Println(ds)

	if ds[0][0].Language != sourceLanguage {
		fmt.Printf("source language is %q, but expected is %q\n", ds[0][0].Language.String(), sourceLanguage.String())
	}

	if ds[0][0].Confidence < 0.9 {
		fmt.Printf("source language detection's confidence for %q is below 90%\n", ds[0][0].Language.String())
	}

	translations, err := translateClient.Translate(ctx,
		[]string{text}, targetLanguage,
		&translate.Options{
			Source: ds[0][0].Language,
			Format: translate.Text,
		})
	if err != nil {
		fmt.Printf("failed to translate: %v\n", err)
		return
	}
	fmt.Println(translations[0].Text)

	//langs, err := translateClient.SupportedLanguages(ctx, language.English)
	//if err != nil {
	//	fmt.Printf("failed to retrieve supported languages: %v\n", err)
	//	return
	//}
	//fmt.Println(langs)
	//
	//for _, language := range langs {
	//	fmt.Println(language.Tag)
	//}

	// Close the client when finished.
	if err := translateClient.Close(); err != nil {
		fmt.Printf("failed to close translate client: %v\n", err)
		return
	}

	pubsubCredentialFile := os.Getenv("SMBE_APPLICATION_CREDENTIALS")
	if pubsubCredentialFile == "" {
		fmt.Printf("SMBE_APPLICATION_CREDENTIALS not set\n")
		return
	}

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(pubsubCredentialFile)

	pubsubClient, err := pubsub.NewClient(ctx, "hybrid-cloud-22365", opt)
	if err != nil {
		fmt.Printf("failed to create new pubsub client: %v\n", err)
		return
	}

	topic := pubsubClient.Topic("smbe_output")
	defer topic.Stop()
	var results []*pubsub.PublishResult
	r := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("hello world"),
	})
	results = append(results, r)
	// Do other work ...
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			fmt.Printf("failed to get pubsub result: %v\n", err)
			return
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
}
