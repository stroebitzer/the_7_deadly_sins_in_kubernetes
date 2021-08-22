#!/bin/bash

export PROJECT_NAME=loodse-training-playground
export REGION=europe-west3
export ZONE=europe-west3-a
export CLUSTER_NAME=hubert-talk
export NETWORK_NAME=$CLUSTER_NAME
export FIREWALL_NAME=$CLUSTER_NAME

set -euxo pipefail

# set gcloud params
gcloud config set project $PROJECT_NAME
gcloud config set compute/region $REGION
gcloud config set compute/zone $ZONE

# delete resources
gcloud beta container clusters delete $CLUSTER_NAME --quiet
gcloud compute firewall-rules delete $FIREWALL_NAME-ingress-gateway --quiet
gcloud compute firewall-rules delete $FIREWALL_NAME-ssh --quiet
gcloud compute firewall-rules delete $FIREWALL_NAME-nodeport --quiet
gcloud compute networks subnets delete $NETWORK_NAME-subnet --quiet
gcloud compute networks delete $NETWORK_NAME --quiet
