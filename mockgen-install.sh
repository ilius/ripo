#!/bin/sh
set -e

MOCKGEN=github.com/ilius/mock
MOCKGEN_REV=

dir="$GOPATH/src/$MOCKGEN/"

if [ -d "$dir" ] ; then
	cd "$dir"
	if [ -n "$MOCKGEN_REV" ] && ! git checkout "$MOCKGEN_REV" ; then
		git fetch
		git checkout "$MOCKGEN_REV"
	fi
else
	git clone https://$MOCKGEN/ "$dir"
	cd "$dir"
	git checkout "$MOCKGEN_REV"
fi

go install ./gomock
go install ./mockgen
