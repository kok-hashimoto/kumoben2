#!/bin/sh
gcloud functions deploy sample-function --allow-unauthenticated --entry-point=EntryPoint --gen2 --region=us-central1 --runtime=go120 --trigger-http
