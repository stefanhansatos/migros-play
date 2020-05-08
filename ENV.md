```bash

export FIREBASE_PROJECT=<project name>
export FIREBASE_URL="https://${FIREBASE_PROJECT}.firebaseio.com"

export FIREBASE_SERVICE_ACCOUNT=<firebase service account>

export FIREBASE_APPLICATION_CREDENTIALS=<json key file of service account> 
export GOOGLE_APPLICATION_CREDENTIALS=<json key file of service account>
# todo: one is enough

export LOCAL_DATA_DIR=<local directory with json data>

gcloud config set account $FIREBASE_SERVICE_ACCOUNT
export ACCESS_TOKEN=$(gcloud config config-helper --format='value(credential.access_token)')
```
