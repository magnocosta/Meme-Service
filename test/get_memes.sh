#!/bin/bash

LAT=$1
LON=$2
QUERY=$3

required_vars=("LAT" "LON" "QUERY")

check_required_vars() {
    for var in "${required_vars[@]}"; do
        if [ -z "${!var}" ]; then
            echo "Error: $var is required and not set or empty."
            exit 1
        fi
    done
}

check_required_vars

curl -s "http://localhost:8080/memes?query=${QUERY}&lat=${LAT}&long=${LON}" \
  -H "Content-Type: application/json" \
  | jq
