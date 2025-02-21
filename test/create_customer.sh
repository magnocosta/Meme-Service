#!/bin/bash

NAME=$1
EMAIL=$2
PAYLOAD=$(cat <<EOF
{
  "email": "${EMAIL}",
  "name": "${NAME}"
}
EOF)

required_vars=("NAME" "EMAIL")

check_required_vars() {
    for var in "${required_vars[@]}"; do
        if [ -z "${!var}" ]; then
            echo "Error: $var is required and not set or empty."
            exit 1
        fi
    done
}

check_required_vars

curl -s "http://localhost:8080/v1/customers" \
  -H "Content-Type: application/json" \
  -d "${PAYLOAD}" \
  | jq
