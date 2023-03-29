#!/bin/sh
gcloud api-gateway api-configs create sample-config --api=sample-api --project=kumoben2-test1 --openapi-spec=spec-v3.yaml
