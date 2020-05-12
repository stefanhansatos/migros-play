We store [global variables](../ENV.md) locally to use them all over the place.


Todo: Consistent naming, refactoring, and commenting
Todo: [One-time initialization](https://cloud.google.com/functions/docs/concepts/go-runtime#one-time_initialization)

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

curl https://$FIREBASE_REGION-$FIREBASE_PROJECT.cloudfunctions.net/appendhttp
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

gcloud pubsub topics publish fb_someData --message "not yet used by Append"

gcloud pubsub subscriptions describe projects/hybrid-cloud-22365/subscriptions/gcf-Append-europe-west1-fb_someData
gcloud pubsub subscriptions delete projects/hybrid-cloud-22365/subscriptions/gcf-Append-europe-west1-fb_someData
```
---

[func Store(ctx context.Context, m PubSubMessage) error](./pubsub_store.go)

```bash
gcloud functions deploy Store --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=fb_someData \
   --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
   --service-account=${FIREBASE_SERVICE_ACCOUNT}
   
gcloud functions call Store --region ${FIREBASE_REGION} --data='{"message": "Hello World!"}'

gcloud pubsub topics publish fb_someData --message "Payload: foo at $(date)"

gcloud pubsub subscriptions describe projects/hybrid-cloud-22365/subscriptions/gcf-Store-europe-west1-fb_someData
gcloud pubsub subscriptions delete projects/hybrid-cloud-22365/subscriptions/gcf-Store-europe-west1-fb_someData 
```
---

[func WrapPayload(ctx context.Context, message Message) error](./pubsub_wrap_payload.go)

```bash
gcloud functions deploy WrapPayload --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=fb_someData \
   --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
   --service-account=${FIREBASE_SERVICE_ACCOUNT}
   
DATA=$(printf '{ "type": "int", "data": 22 }'|base64) && gcloud functions call WrapPayload --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'

gcloud pubsub topics publish fb_someData --message '{ "type": "int", "data": 22 }'
```
---

#### BigQuery

[func Http_Query(w http.ResponseWriter, r *http.Request) error](./http_bq_query.go)

```bash
gcloud functions deploy query --region ${FIREBASE_REGION} --entry-point Http_Query --runtime go111 --trigger-http
gcloud functions call query  --region ${FIREBASE_REGION} --data '{}'

curl https://$FIREBASE_REGION-$FIREBASE_PROJECT.cloudfunctions.net/query
```
---

[BqQuery(ctx context.Context, m interface{}) error](./pubsub_bq_query.go)

```bash
gcloud functions deploy BqQuery --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=fb_someData
gcloud functions call BqQuery --region ${FIREBASE_REGION} --data '{}'

gcloud pubsub topics publish fb_someData --message "Payload: foo at $(date)"
```
---

#### Event

[func HelloGCS(ctx context.Context, e GCSEvent) error](./event_storage.go)

```bash
gcloud functions deploy HelloGCS --region ${FIREBASE_REGION} --runtime go111 \
   --trigger-event="providers/cloud.storage/eventTypes/object.change" --trigger-resource="hybrid-cloud-22365.appspot.com" \
   --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
   --service-account=${FIREBASE_SERVICE_ACCOUNT}

gcloud functions deploy HelloGCSInfo --region ${FIREBASE_REGION} --runtime go111 \
   --trigger-event="providers/cloud.storage/eventTypes/object.change" --trigger-resource="hybrid-cloud-22365.appspot.com" \
   --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
   --service-account=${FIREBASE_SERVICE_ACCOUNT}

gsutil cp ../data/notes.txt gs://hybrid-cloud-22365.appspot.com/notes.txt
gsutil rm gs://hybrid-cloud-22365.appspot.com/notes.txt
gsutil cp ../data/notes.txt gs://hybrid-cloud-22365.appspot.com/notes.txt
gsutil cp ../data/notes.txt gs://hybrid-cloud-22365.appspot.com/notes.txt
gsutil rm gs://hybrid-cloud-22365.appspot.com/notes.txt

gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="HelloGCS" resource.labels.region="europe-west1" severity=DEFAULT'  \
  | grep textPayload | head -5

gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="HelloGCSInfo" resource.labels.region="europe-west1" severity=DEFAULT'  \
  | grep textPayload | head -35
```
---

[func HelloRTDB(ctx context.Context, e RTDBEvent) error](./event_database.go)

```bash
gcloud functions deploy HelloRTDB --region ${FIREBASE_REGION} --runtime go111 \
   --trigger-event="providers/google.firebase.database/eventTypes/ref.write" --trigger-resource="projects/_/instances/hybrid-cloud-22365/refs/someData/list" \
   --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
   --service-account=${FIREBASE_SERVICE_ACCOUNT}
   
DATA=$(printf '{ "type": "int", "data": 22 }'|base64) && gcloud functions call WrapPayload --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'

gcloud logging read 'resource.type="cloud_function" resource.labels.function_name="HelloRTDB" resource.labels.region="europe-west1" severity=DEFAULT'  \
  | grep textPayload | head -35   
```



#### Raw Functions as Standalone

After deleting the Pub/Sub subscription can only be called by `gcloud functions call` (or programmatically).

[func RawWrapPayload(ctx context.Context, message Message) error](./raw_wrap_payload.go)
```bash
gcloud functions deploy RawWrapPayload --region ${FIREBASE_REGION} --runtime go111 --trigger-topic=fb_someData \
   --set-env-vars=FIREBASE_PROJECT=${FIREBASE_PROJECT},FIREBASE_URL=${FIREBASE_URL} \
   --service-account=${FIREBASE_SERVICE_ACCOUNT}
   
gcloud pubsub subscriptions delete projects/hybrid-cloud-22365/subscriptions/gcf-RawWrapPayload-europe-west1-fb_someData
   
DATA=$(printf '{ "type": "int", "key": "count", "value": 4 }'|base64) && \
  gcloud functions call RawWrapPayload --region ${FIREBASE_REGION} --data '{"data":"'$DATA'"}'
```
---

