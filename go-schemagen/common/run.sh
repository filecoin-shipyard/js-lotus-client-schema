#! /bin/bash

go run schemagen.go ../lotus | jq -M
