#!/bin/sh
gcloud api-gateway gateways delete sample-gateway --location=us-central1
gcloud api-gateway api-configs delete sample-config --api=sample-api
