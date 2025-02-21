#!/bin/bash

CUSTOMER_ID=$1
TOKENS=$2
PAYLOAD=$(cat <<EOF
{
  "customer_id": "${CUSTOMER_ID}",
  "tokens": ${TOKENS}
}
EOF)

required_vars=("CUSTOMER_ID" "TOKENS")

check_required_vars() {
    for var in "${required_vars[@]}"; do
        if [ -z "${!var}" ]; then
            echo "Error: $var is required and not set or empty."
            exit 1
        fi
    done
}

check_required_vars

curl -s "http://localhost:8080/v1/tokens" \
  -H "Content-Type: application/json" \
  -d "${PAYLOAD}" \
  | jq
