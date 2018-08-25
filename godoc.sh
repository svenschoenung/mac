#!/usr/bin/env bash

set -m
godoc -http=":8181" &
open "http://localhost:8181/pkg/${PWD#$GOPATH/src}"
fg 1
