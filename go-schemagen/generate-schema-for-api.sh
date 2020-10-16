#! /bin/bash

API=$1
if [ -z "$API" ]; then
  echo Specify API: Common, FullNode, StorageMiner, GatewayAPI, WalletAPI, WorkerAPI
  exit 1
fi
go run . $API | jq -M
