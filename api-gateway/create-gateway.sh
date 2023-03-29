#!/bin/sh
gcloud api-gateway gateways create sample-gateway --api=sample-api --api-config=sample-config --location=us-central1
