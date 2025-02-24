#!/bin/bash

CUSTOMER_ID=$1

curl -s "http://localhost:8080/v1/customers/${CUSTOMER_ID}/tokens" \
  -H "Content-Type: application/json" \
  | jq
