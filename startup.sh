#! /bin/sh

# Small script that installs gin (build helper) and dep (dependency manager)
go get github.com/codegangsta/gin
go get -u github.com/golang/dep/cmd/dep
/go/bin/dep ensure
/go/bin/gin -a 8080 -d /go/src/github.com/willdobbins/notes/cmd/notes/ --path /go/src/github.com/willdobbins/notes run notes.go
