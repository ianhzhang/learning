
go run server_upload.go

1. http client:
python3 -m http.server
http://localhost:8000

2. curl client:
curl -F 'myFile=@index.html' http://localhost:8080/upload

3. go client:
go run client_upload.go 

4. python client:
python3 client_upload.py

5. Docker server:
go mod init upload_server
go build -v -o server_upload

dockerbuild
dockerrun
dockerbash my-ubuntu

docker logs my-ubuntu
