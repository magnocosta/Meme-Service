#!/bin/bash

LAT=$1
LON=$2
QUERY=$3

curl -s "http://localhost:8080/v1/memes?query=${QUERY}&lat=${LAT}&lon=${LON}" \
  -H "Content-Type: application/json" \
  | jq
