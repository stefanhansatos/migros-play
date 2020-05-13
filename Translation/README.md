```bash
go get -u cloud.google.com/go/translate
go mod vendor
```

```bash
gcloud services list --available --filter="name ~ .*translate.googleapis.com"
gcloud services list --filter="name ~ .*translate.googleapis.com"

gcloud services enable translate.googleapis.com

gcloud services list --filter="name ~ .*translate.googleapis.com"
```

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

```bash
go run main.go
``` 