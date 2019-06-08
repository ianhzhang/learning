cd src/gcptest

go mod init gcptest
go mod tidy

gcloud functions deploy IanFirestoreTest --runtime go111 --trigger-http