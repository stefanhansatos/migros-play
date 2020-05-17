#### Overview of Serverless Mobile Back End

```
[mobile] --> "https://<...>/translate", i.e. func SmbeHTTP
```

```
func SmbeHTTP ->  pubsub topic "smbe_input"
```

```
pubsub topic "smbe_input" --> func SmbeTranslationQueryLoad --> realtime db "https://<...>/translation/queries"

pubsub topic "smbe_input" --> func SmbeTranslate
```

```
func SmbeTranslate <--> service Translation API

func SmbeTranslate  --> pubsub topic "smbe_output"
```

```
pubsub topic "smbe_output" --> function SmbeTranslationLoad --> realtime db "https://<...>/translation/results"

pubsub topic "smbe_output" --> function SmbeFileStore --> storage "gs://smbe-..."
```

```
realtime db "https://<...>/translation/results" -> function "SendOutput" --> [mobile]
```

#### Environment

We store [global variables](../ENV.md) locally to use them all over the place.

```bash
export SHORT_NAME="smbe" # serverless-mobile-back-end
export RTDB_URL="https://${SHORT_NAME}.firebaseio.com"
```



Create service account, 
```bash
gcloud iam service-accounts create ${SMBE_NAME} \
    --description="Service account to publish pubsub messages" \
    --display-name="${SMBE_NAME}"
    
gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:${SMBE_NAME}@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/pubsub.publisher
    
gcloud iam service-accounts keys create ${LOCAL_CREDENTIALS_DIR}/${GCP_PROJECT}-${SMBE_NAME}.json \
  --iam-account ${SMBE_NAME}@${GCP_PROJECT}.iam.gserviceaccount.com
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

#### Functions 

[func SmbeHTTP(ctx context.Context, message Message) error](./http-frontend/functions.go)

```bash
cd ../http-frontend

gcloud functions deploy translate --region ${FIREBASE_REGION}  --entry-point SmbeHTTP --runtime go111 --trigger-http \
    --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},SMBE_PUBSUB_TOPIC_IN=${SHORT_NAME}_input \
    --service-account=${HTTP_SERVICE_ACCOUNT}

# Public accessable
gcloud alpha functions add-iam-policy-binding translate --region=europe-west1 --member=allUsers --role=roles/cloudfunctions.invoker"

export ACCESS_TOKEN=$(gcloud config config-helper --format='value(credential.access_token)')

DATA=$(printf '{ "text": "Today is Monday", "sourceLanguage": "en",  "targetLanguage": "fr"}'|base64) && curl -X POST "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/translate" \
  -d "'$DATA'"
  \--data '{"data":"'$DATA'"}'
  

curl -X POST "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/translate" \
  -d '{ "text":"Hallo alle zusammen. Wie geht es?","sourceLanguage":"de", "targetLanguage": "fr" }'


gsutil cat gs://hybrid-cloud-22365.appspot.com/beab10c6-deee-4843-9757-719566214526
```
---

[func SmbeTranslationQueryLoad(ctx context.Context, message Message) error](./realtime-db/functions.go)

```bash
gcloud functions deploy SmbeTranslationQueryLoad --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=${SHORT_NAME}_input \
  --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},RTDB_URL=${RTDB_URL} \
  --service-account=${FIREBASE_SERVICE_ACCOUNT}

gcloud functions call SmbeTranslationQueryLoad --region ${FIREBASE_REGION} --data '{}'

DATA=$(printf '{ "text": "SmbeTranslationQueryLoad: Today is Monday", "sourceLanguage": "en",  "targetLanguage": "fr"}'|base64) && gcloud functions call SmbeTranslationQueryLoad --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'


gcloud pubsub topics publish ${SHORT_NAME}_input --message '{ "text": "2: Tommorow is Tuesday", "sourceLanguage": "en",  "targetLanguage": "fr"}'

gcloud pubsub subscriptions describe projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationQueryLoad-europe-west1-${SHORT_NAME}_input
gcloud pubsub subscriptions delete projects/hybrid-cloud-22365/subscriptions/gcf-SmbeTranslationQueryLoad-europe-west1-${SHORT_NAME}_input
```
---


[func SmbeTranslate(ctx context.Context, message Message) error](./translation/functions.go)

```bash
cd translation

gcloud functions deploy SmbeTranslate --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=${SHORT_NAME}_input \
  --set-env-vars=RTDB_URL=${RTDB_URL} \
  --service-account=${FIREBASE_SERVICE_ACCOUNT}


DATA=$(printf '{ "text": "4: Today is not Monday", "sourceLanguage": "en",  "targetLanguage": "fr"}'|base64) && \
  gcloud functions call SmbeTranslate --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'
  
gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="SmbeTranslate" resource.labels.region="europe-west1" severity=DEFAULT' \
   --format json | head -35


gcloud pubsub topics publish ${SHORT_NAME}_input --message '{ "text": "Tommorow is Tuesday", "sourceLanguage": "en",  "targetLanguage": "fr"}'

gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="SmbeTranslate" resource.labels.region="europe-west1" severity=DEFAULT' \
   --format json | less
```

---
[func SmbeTranslationLoad(ctx context.Context, message Message) error]()

```bash
gcloud functions deploy SmbeTranslationLoad --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=${SHORT_NAME}_output \
  --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},RTDB_URL=${RTDB_URL} \
  --service-account=${FIREBASE_SERVICE_ACCOUNT}

DATA=$(printf '{ "translationQuery": {"text": "4: Today is not Monday", "sourceLanguage": "en",  "targetLanguage": "fr"},"translatedText": "tranlated", ["lalal", "lllulu"]}'|base64) && \
  gcloud functions call SmbeTranslationLoad --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'
  
  
execution_id=$(DATA=$(printf '{ "translationQuery": {"text": "4: Today is not Monday", "sourceLanguage": "en",  "targetLanguage": "fr"},"translatedText": "tranlated", "translationErrors": ["lalal", "lllulu"] }'|base64) \
&& gcloud functions call SmbeTranslationLoad --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}' |awk -F": " '{ print $2 }') \
&& echo 'gcloud logging read "labels.execution_id=$execution_id severity=DEFAULT" | grep textPayload'
  
gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="SmbeTranslationLoad" resource.labels.region="europe-west1" severity=DEFAULT labels.execution_id="0ubpxrdwm3u3"' \

```




---
[func SmbeFileStore(ctx context.Context, message Message) error](./storage/functions.go)
```bash
gcloud functions deploy SmbeFileStore --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=${SHORT_NAME}_output \
  --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_BUCKET_URL=${FIREBASE_BUCKET_URL} \
  --service-account=${STORAGE_SERVICE_ACCOUNT}

DATA=$(printf '{ "translationQuery": {"text": "4: Today is not Monday", "sourceLanguage": "en",  "targetLanguage": "fr"},"translatedText": "tranlated", ["lalal", "lllulu"]}'|base64) && \
  gcloud functions call SmbeFileStore --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'
  
gsutil ls -l gs://hybrid-cloud-22365.appspot.com/
gsutil ls -L gs://hybrid-cloud-22365.appspot.com/directcall
gsutil cat gs://hybrid-cloud-22365.appspot.com/directcall


gcloud pubsub topics publish ${SHORT_NAME}_output \
  --message '{ "translationQuery": {"text": "4: Today is not Monday", "sourceLanguage": "en",  "targetLanguage": "fr"},"translatedText": "tranlated", ["lalal", "lllulu"]}'


```