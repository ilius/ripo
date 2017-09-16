#!/bin/bash

go get -u github.com/golang/mock/gomock
go get -u github.com/golang/mock/mockgen

rm mock_request.go 2>/dev/null

mockgen \
	-package restpc \
	-source request.go \
	-destination mock_request.go \
	|| exit $?
