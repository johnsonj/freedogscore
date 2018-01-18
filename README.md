# FreeDogScore

This repo contains the open source rewrite of [FreeDogScore.com](https://freedogscore.com). The goal is to replace the Rails/Postgres application with [Cloud Functions](https://cloud.google.com/functions/) that generate the static website.

# Development

## Dependencies
- Linux environment
- [Golang 1.9+](https://golang.org/)
- [Node 9.4+](https://nodejs.org/en/) (only needed for deployment)
- [gcloud sdk](https://cloud.google.com/sdk/downloads)
- [Google Cloud Project](https://cloud.google.com)
- [Terraform 0.11.0+](https://terraform.io) (only needed for deployment)

## Running Locally
```bash
go get github.com/johnsonj/freedogscore/web
cd $(go env GOPATH)/src/github.com/johnsonj/freedogscore/web
go run main.go
```

## Deploying as a Cloud Function

```bash
cd $(go env GOPATH)/src/github.com/johnsonj/freedogscore/functions
make deploy
```
