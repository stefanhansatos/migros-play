We store [global variables](../ENV.md) locally to use them all over the place.

#### Go Client Library

Create service account, 
```bash
gcloud iam service-accounts create ${FIREBASE_BUCKET_NAME} \
    --description="Service account to invoke client library for Storage" \
    --display-name="${FIREBASE_BUCKET_NAME}"
    
gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:${FIREBASE_BUCKET_NAME}@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/storage.admin
    
gcloud iam service-accounts keys create ${LOCAL_CREDENTIALS_DIR}/${GCP_PROJECT}-${FIREBASE_BUCKET_NAME}.json \
  --iam-account ${FIREBASE_BUCKET_NAME}@${GCP_PROJECT}.iam.gserviceaccount.com
  
go get -u cloud.google.com/go/storage
go mod vendor
```
