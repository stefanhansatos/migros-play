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
go run main.go
```