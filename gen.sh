#!/bin/bash
set -e
# set -o xtrace

./mockgen-install.sh

rm mock_*.go 2>/dev/null || true

RIPO=github.com/ilius/ripo

function fix_mock_file() {
	sed -i 's|package mock_ripo|package ripo|g' "$1"
}

function gen_internal_mock() {
	local interface_names="$1"
	local output_path="$2"
	src="$($GOPATH/bin/mockgen -self_package $RIPO $RIPO $interface_names)" || return $?
	if [ -n "$src" ] ; then
		echo "$src" > "$output_path"
		fix_mock_file "$output_path"
	fi
}

$GOPATH/bin/mockgen -package ripo -destination mock_readcloser.go io ReadCloser  || exit $?

gen_internal_mock Request,ExtendedRequest mock_request.go

gen_internal_mock SmallT mock_t.go 

