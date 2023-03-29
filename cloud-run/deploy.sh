#!/bin/sh
PROJECT_ID=`gcloud config get project`
gcloud run deploy sample-run --allow-unauthenticated --execution-environment=gen2 --image=us-central1-docker.pkg.dev/${PROJECT_ID}/registry/app:latest --region=us-central1
