## Golang and MUX router REST API

## Install Golang - Linux

```bash
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz

vim ~/.profile

(append :/usr/local/go/bin to PATH)
```

## Go Modules - Initialize the module for the app (from root dir)

```bash
go mod init gitlab.com/pragmaticreviews/golang-mux-api
```

## Go Modules - Download external dependencies

```bash
go mod download  
```

# The go command uses the go.sum file to ensure that future downloads of these modules retrieve the same bits as the first download (it keeps an initial hash)
#Both go.mod and go.sum should be checked into version control.

## Export Environment variable (Firestore)

```bash
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private-key.json'
```

## How to get the private key JSON file:
## From the Firebase Console: Project Overview -> Project Settings -> Service Accounts -> Generate new private key

## Build

```bash
go build
```
## Test (specific test)

```bash
go test -run NameOfTest
```

## Test (all the tests within the service folder)

```bash
go test service/*.go
```

## Run

```bash
go run .
```

```bash
go run *.go
```

## Docker

```bash
docker build -t golang-api .
```

```bash
docker run -p 8000:8000 golang-api
```