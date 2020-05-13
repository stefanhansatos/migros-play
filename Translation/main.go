package main

import (
	"context"
	"fmt"

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
	client, err := translate.NewClient(ctx)
	if err != nil {
		fmt.Printf("failed to create new client: %v\n", err)
		return
	}

	// Use the client.
	ds, err := client.DetectLanguage(ctx, []string{text})
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

	translations, err := client.Translate(ctx,
		[]string{text}, targetLanguage,
		&translate.Options{
			Source: ds[0][0].Language,
			Format: translate.Text,
		})
	if err != nil {
		// TODO: handle error.
	}
	fmt.Println(translations[0].Text)

	// Close the client when finished.
	if err := client.Close(); err != nil {
		// TODO: handle error.
	}
}
