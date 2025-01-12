# Klaviyo to BigQuery

This project is an application that retrieves data from the Klaviyo API and stores it in BigQuery. Currently you can import event, profile, and metric data. All data is newly created and existing data is not updated.

## Requirements

- Go 1.22 and later
- Klaviyo API Keys
- BigQuery project
- Google Cloud SDK
- Google Cloud project
- Enable Google Cloud BigQuery API

## Setup 

```sh
git clone https://github.com/yourusername/klaviyo-to-bigquery.git
cd klaviyo-to-bigquery
go mod tidy
go build -o klaviyo-to-bigquery
```

## Configuration Files
```
{
  "KEY": "KLAVIYO_API_KEY",
  "CLIENT_SECRET_FILE": google_client_secret.json,
  "SERVICE_ACCOUNT_FILE": google_service_account.json,
  "PROJECT_ID": "YOUR_PROJECT_ID",
  "PROPERTY_ID": "YOUR_PROPERTY_ID",
  "FETCH_TO_DATE": "2024-10-04T05:00:00Z",
  "DATASET_ID": "klaviyo_api",
  "TABLE_PREFIX": "klaviyo_"
}

```


## Usage

```sh
./klaviyo-to-bigquery [subcommand] --api-key pk_xxxxxxxxx   --fetch-to-date 2024-11-17T00:00:00Z --config ./config.yaml
```

### Sub Commands

- getEvents: Gets Klaviyo event data.
- getProfiles: Gets Klaviyo profile data.
- getMetrics: Gets Klaviyo metrics data.


### Run

```sh
go run main.go <command> --config=<Config file>
```

## TODO
- [x] Getting data from Klaviyo APIs
- [x] Storing data in BigQuery
- [ ] Writing test code
- [ ] Writing documentation
- [ ] Adding logs
- [ ] Write a Dockerfile

## License
This project is released under the Apache-2.0 license. For more information, see [LICENSE](LICENSE).
