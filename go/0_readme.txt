export GOPATH="$HOME/goproj/proj"

cd $GOPATH/src
go get github.com/gorilla/mux

This will create
proj/pkg/
proj/src/github.com/


In proj/src:
go run main.go



=========================
script
directory structure:
  proj
  proj/src
  proj/src/main.go
=========================
cd proj

#1. in proj directory, set proj directory as GOPATH
export GOPATH=`pwd`
echo $GOPATH

#2. in proj directory, install 3rd party software
go get github.com/gorilla/mux

#3 in src directory run 
cd src
go run main.go


read key value pair from received json