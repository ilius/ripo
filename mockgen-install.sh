#!/bin/sh
set -e

MOCKGEN_REV=v1.2.0

dir="$GOPATH/src/github.com/golang/mock/"

if [ -d "$dir" ] ; then
	cd "$dir"
	if ! git checkout "$MOCKGEN_REV" ; then
		git fetch
		git checkout "$MOCKGEN_REV"
	fi
else
	git clone https://github.com/golang/mock/ "$dir"
	cd "$dir"
	git checkout "$MOCKGEN_REV"
fi

go install ./gomock
go install ./mockgen
