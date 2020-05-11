Project

```bash

export GCP_PROJECT=<project name>
export FIREBASE_PROJECT=<project name>
# todo: one is enough

export FIREBASE_REGION=<project region>

export FIREBASE_SERVICE_ACCOUNT=<firebase service account>

export FIREBASE_APPLICATION_CREDENTIALS=<json key file of service account> 
export GOOGLE_APPLICATION_CREDENTIALS=<json key file of service account>
# todo: one is enough

```

Local
```bash
export LOCAL_DATA_DIR=<local data directory>
export LOCAL_CREDENTIALS_DIR=<local credentials directory>
```

Firebase
```bash
export FIREBASE_URL="https://${FIREBASE_PROJECT}.firebaseio.com"
```

REST
```bash
gcloud config set account $FIREBASE_SERVICE_ACCOUNT
export ACCESS_TOKEN=$(gcloud config config-helper --format='value(credential.access_token)')
```

BigQuery
```bash
export BQ_SA_NAME=<service account name>
export BIGQUERY_APPLICATION_CREDENTIALS="${LOCAL_CREDENTIALS_DIR}/${GCP_PROJECT}-${BQ_SA_NAME}.json"
```

Storage
```bash
export FIREBASE_BUCKET_NAME=<name of the storage bucket>
export FIREBASE_BUCKET_URL="${FIREBASE_BUCKET_NAME}.appspot.com"
export STORAGE_APPLICATION_CREDENTIALS="${LOCAL_CREDENTIALS_DIR}/${GCP_PROJECT}-${FIREBASE_BUCKET_NAME}.json"
```
