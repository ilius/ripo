#!/bin/bash

go get -u github.com/golang/mock/gomock
go get -u github.com/golang/mock/mockgen

rm mock_request.go 2>/dev/null

$GOPATH/bin/mockgen \
	-package ripo \
	-source request.go \
	-destination mock_request.go \
	|| exit $?

$GOPATH/bin/mockgen -package ripo \
	-destination mock_readcloser.go \
	io ReadCloser  || exit $?
