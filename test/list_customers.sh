#!/bin/bash

curl -s "http://localhost:8080/v1/customers" \
  -H "Content-Type: application/json" \
  | jq
