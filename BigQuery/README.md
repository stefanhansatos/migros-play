We store [global variables](../ENV.md) locally to use them all over the place.

#### bq

```bash
gcloud services enable bigquery.googleapis.com

bq show
```

```bash
bq help

bq show --help
bq show
bq show bigquery-public-data:samples.shakespeare
bq show bigquery-public-data:samples

bq ls --help
bq ls
bq ls bigquery-public-data:samples
```

```bash
bq query --use_legacy_sql=false \
'SELECT
   word,
   SUM(word_count) AS count
 FROM
   `bigquery-public-data`.samples.shakespeare
 WHERE
   word LIKE "%raisin%"
 GROUP BY
   word'
```

Download example data names.zip from [here](http://www.ssa.gov/OACT/babynames/names.zip) to the [local data directory](../ENV.md).
```bash
cd ${LOCAL_DATA_DIR}
unzip names.zip
```

```bash
bq show
bq ls

bq mk babynames 
bq ls
bq ls babynames
bq show babynames

cd ${LOCAL_DATA_DIR}
head yob2010.txt

bq load babynames.names2010 yob2010.txt name:string,gender:string,count:integer
bq show babynames.names2010
bq query 'SELECT count(*) from babynames.names2010'

bq query "SELECT name,count FROM babynames.names2010 WHERE gender = 'F' ORDER BY count DESC LIMIT 5"
```

#### Go Client Library

Create service account, 
```bash
gcloud iam service-accounts create ${BQ_SA_NAME} \
    --description="Service account to invoke client library for BigQuery" \
    --display-name="${BQ_SA_NAME}"
    
gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
  --member serviceAccount:${BQ_SA_NAME}@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/bigquery.admin
    
gcloud iam service-accounts keys create ${LOCAL_CREDENTIALS_DIR}/${GCP_PROJECT}-${BQ_SA_NAME}.json \
  --iam-account ${BQ_SA_NAME}@${GCP_PROJECT}.iam.gserviceaccount.com
```
