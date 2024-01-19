#!/bin/sh
mkdir -p /go/src/main
cd /go/src/main
swag init
air -c .air.toml