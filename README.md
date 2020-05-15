# migros-play
Creative Learning around GCP, Firebase, and more

#### [Firebase Database REST API](https://firebase.google.com/docs/reference/rest/database)

- [README](./REST/README.md)

#### [Firebase Admin SDK](https://firebase.google.com/docs/admin/setup)

- [Go](https://github.com/firebase/firebase-admin-go)
- [JavaScript](https://github.com/firebase/firebase-admin-node)


#### [Firebase CLI](https://firebase.google.com/docs/cli)

- [firebase-tools](https://github.com/firebase/firebase-tools)
- [README](./CLI/README.md)

#### [Cloud Functions](https://firebase.google.com/docs/functions/functions-and-firebase)

- [Cloud Functions](https://cloud.google.com/functions/docs)
- [Cloud Functions for Firebase](https://firebase.google.com/docs/functions)
- [Admin SDK DB](https://firebase.google.com/docs/database/admin/start#go)
- [README](./functions/README.md)

#### [BigQuery](https://cloud.google.com/bigquery/docs)

- [README](./BigQuery/README.md)
- [Go Package](https://pkg.go.dev/cloud.google.com/go/bigquery?tab=doc)
- [Go Samples](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/bigquery)

#### [Storage](https://cloud.google.com/storage/docs)

- [README](./Storage/README.md)
- [Go Package](https://pkg.go.dev/cloud.google.com/go/storage?tab=doc)
- [Go Samples](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/storage)

---

[Event provider](https://cloud.google.com/functions/docs/concepts/events-triggers) 
trigger cloud functions in the background:

- Cloud Pub/Sub
- Firebase Realtime Database
- Firestore Database
- Cloud Storage
- Firebase Analytics Events

Cloud Functions can be called from [Cloud Scheduler](https://cloud.google.com/scheduler/docs/tut-pub-sub) via Cloud Pub/Sub. 
Most commonly, from outside the cloud Cloud Functions are called via [HTTP requests](https://cloud.google.com/functions/docs/calling/http).

Cloud Functions written in Go use the 
[Admin SDK](https://firebase.google.com/docs/admin/setup#prerequisites)
 for certain [Firebase features](https://pkg.go.dev/firebase.google.com/go?tab=doc):
 
- [Firebase Cloud Messaging](https://pkg.go.dev/firebase.google.com/go@v3.12.1+incompatible/messaging?tab=doc)
- [Firebase Realtime Database](https://pkg.go.dev/firebase.google.com/go@v3.12.1+incompatible/db?tab=doc)
- [Firebase Storage](https://pkg.go.dev/firebase.google.com/go@v3.12.1+incompatible/storage?tab=doc)

There are many services more, which are not related directly to or not accessed via Firebase:
- [Firestore](https://pkg.go.dev/cloud.google.com/go/firestore?tab=doc)
- [BigQuery](https://pkg.go.dev/cloud.google.com/go/bigquery?tab=doc)
- [Translate](https://pkg.go.dev/cloud.google.com/go@v0.57.0/translate?tab=doc)
- [Pub/Sub](https://pkg.go.dev/cloud.google.com/go/pubsub?tab=doc)
- [and more ...](https://pkg.go.dev/cloud.google.com/go?tab=subdirectories)




