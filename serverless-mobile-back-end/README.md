#### Overview of Serverless Mobile Back End

[mobile] --> pubsub "Input"

pubsub "Input" --> function "LoadInput" --> realtime db "Input"
pubsub "Input" --> function "Translate" --> pubsub "Output"


pubsub "Output" --> function "LoadOutput" --> realtime db "Output"
pubsub "Output" --> function "LoadOutput" --> storage "Output"

realtime db "Output" --> function "SendOutput" --> [mobile]

#### Environment

We store [global variables](../ENV.md) locally to use them all over the place.

```bash
export SHORT_NAME="smbe" # serverless-mobile-back-end
export RTDB_URL="https://${SHORT_NAME}.firebaseio.com"

```

#### Pub/Sub

```bash
gcloud pubsub topics create ${SHORT_NAME}_input
gcloud pubsub topics create ${SHORT_NAME}_output
```

#### Realtime Database

```bash
firebase database:instances:create ${SHORT_NAME}
```

#### Storage

```bash
gsutil mb gs://${SHORT_NAME}-hybrid-cloud-22365/
```

#### Function "LoadInput"


[func SmbeTranslationQueryLoad(ctx context.Context, message Message) error](./functions.go)

```bash
gcloud functions deploy SmbeTranslationQueryLoad --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=${SHORT_NAME}_input \
  --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},RTDB_URL=${RTDB_URL} \
  --service-account=${FIREBASE_SERVICE_ACCOUNT}

gcloud functions call SmbeTranslationQueryLoad --region ${FIREBASE_REGION} --data '{}'

DATA=$(printf '{ "text": "Today is Monday", "sourceLanguage": "en",  "targetLanguage": "fr"}'|base64) && gcloud functions call SmbeTranslationQueryLoad --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'


gcloud pubsub topics publish ${SHORT_NAME}_input --message '{ "text": "Today is Tuesday", "sourceLanguage": "en",  "targetLanguage": "fr"}'

gcloud pubsub subscriptions describe projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationQueryLoad-europe-west1-${SHORT_NAME}_input
gcloud pubsub subscriptions delete projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationQueryLoad-europe-west1-${SHORT_NAME}_input
```
---


