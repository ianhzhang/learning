alias gopath="export GOPATH=\`pwd\`"

cd proj
gopath

go install 1_hello_app
go install 2_web_hello_app
go install 3_mux_hello_app

4_json_proj:
cd src/4_json_proj
go run main.go
go run my1.go
go run my2.go