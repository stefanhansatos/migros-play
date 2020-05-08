```bash

export FIREBASE_PROJECT=<project name>
export FIREBASE_URL="https://${FIREBASE_PROJECT}.firebaseio.com"

export FIREBASE_APPLICATION_CREDENTIALS=<json key file of service account> 
export GOOGLE_APPLICATION_CREDENTIALS=<json key file of service account>
# todo: one is enough

export LOCAL_DATA_DIR=<local directory with json data>

export ACCESS_TOKEN=<token for curl bearer>
```


Get the access token.

```bash
gcloud auth list
gcloud config set account `FIREBASE ACCOUNT`

gcloud config config-helper --format='value(credential.access_token)'

export ACCESS_TOKEN=$(gcloud config config-helper --format='value(credential.access_token)')

```