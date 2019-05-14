local test:

gopath
go get all
go build localserver
./localserver

Deploy to google function:
cd src/gcp_functions
gcloud functions deploy Hello2Get --runtime go111 --trigger-http

cd my2
gopath

cd src/gcp_functions/
export GO111MODULE=on
go mod init gcp_functions
go mod tidy
gcloud functions deploy Entry --runtime go111 --trigger-http
