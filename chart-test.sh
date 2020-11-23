#!/bin/bash

# Install the current version of the Helm chart
# to a Kind cluster for testing.

# Set up an empty kubernetes cluster
ctlptl create cluster kind

# Add the repo to Helm
helm repo add tic-tac-toverengineered \
    https://theothertomelliott.github.io/tic-tac-toverengineered/
helm repo update
helm search repo tic-tac-toverengineered

# Install the latest version of the chart
helm install \
    --create-namespace --namespace tictactoe \
    tic-tac-toverengineered/tic-tac-toe --generate-name

kubectl wait --for=condition=available --timeout=60s --namespace tictactoe --all deployments

# Forward port from web service to allow testing, will block
kubectl port-forward --namespace tictactoe deployment/web 8080:8080

# Clean up once done testing
ctlptl delete cluster kind-kind