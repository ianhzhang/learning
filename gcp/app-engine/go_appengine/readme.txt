gopath

cd src/project
go get google.golang.org/appengine
(this will downloadd the package to $GOPATH/pkg, $GOPATH/src)

gcloud app deploy


app.yaml should not be in the same directory as $GOPATH