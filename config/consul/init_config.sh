#!/bin/sh

CONSUL_HOST=$1
echo "Init consul configs..."

for json in $(ls orgs|grep json); do
  curl \
  --request PUT \
  --data @./orgs/"$json" \
  "$CONSUL_HOST"/v1/kv/configs/orgs/$(echo "$json" | cut -d '.' -f 1)

done
