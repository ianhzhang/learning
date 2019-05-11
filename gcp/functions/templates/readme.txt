local test:

gopath
go get all
go build localserver
./localserver

Deploy to google function:
cd src/gcp_functions
gcloud functions deploy Hello2Get --runtime go111 --trigger-http