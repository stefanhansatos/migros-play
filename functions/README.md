We store [global variables](../ENV.md) locally to use them all over the place.


Todo: Consistent naming and deploy using 
- `--service-account=SERVICE_ACCOUNT` 
- `--env-vars-file=FILE_PATH`
- `--allow-unauthenticated=false`

#### HTTP


[func List(w http.ResponseWriter, r *http.Request)](./http_query.go)


```bash
gcloud functions deploy list --region ${FIREBASE_REGION} --entry-point List --runtime go111 --trigger-http
gcloud functions call list  --region ${FIREBASE_REGION} --data '{}'

curl https://${FIREBASE_REGION}-${FIREBASE_PROJECT}.cloudfunctions.net/list
```
---

[func AppendHttp(w http.ResponseWriter, r *http.Request)](./http_append.go)

```bash
gcloud functions deploy appendhttp --region ${FIREBASE_REGION} --entry-point AppendHttp --runtime go111 --trigger-http
gcloud functions call appendhttp  --region ${FIREBASE_REGION} --data '{}'

curl https://${FIREBASE_REGION}-${FIREBASE_PROJECT}.cloudfunctions.net/appendhttp
```
---

#### PUB/SUB

```bash
gcloud pubsub topics create fb_someData
```
---
[func Append(ctx context.Context, r interface{}) error](./pubsub_append.go)

```bash
gcloud functions deploy Append --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=fb_someData \
  --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
  --service-account=${FIREBASE_SERVICE_ACCOUNT}

gcloud functions call Append --region ${FIREBASE_REGION} --data '{}'

gcloud pubsub topics publish fb_someData --message "Payload: foo at $(date)"
```
---

[func Store(ctx context.Context, m PubSubMessage)](./pubsub_store.go)

```bash
gcloud functions deploy Store --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=fb_someData
gcloud functions call Store --region ${FIREBASE_REGION} --data '{"Some payload"}'

gcloud pubsub topics publish fb_someData --message "Some payload"
```
---
