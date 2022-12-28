# tic-tac-toverengineered

[![Release](https://github.com/theothertomelliott/tic-tac-toverengineered/actions/workflows/ci.yaml/badge.svg)](https://github.com/theothertomelliott/tic-tac-toverengineered/actions/workflows/ci.yaml)
[![Tests](https://github.com/theothertomelliott/tic-tac-toverengineered/actions/workflows/test.yaml/badge.svg)](https://github.com/theothertomelliott/tic-tac-toverengineered/actions/workflows/test.yaml)
[![Maintainability](https://api.codeclimate.com/v1/badges/81f71579c34f617680ce/maintainability)](https://codeclimate.com/github/theothertomelliott/tic-tac-toverengineered/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/81f71579c34f617680ce/test_coverage)](https://codeclimate.com/github/theothertomelliott/tic-tac-toverengineered/test_coverage)

A simple Tic Tac Toe game build with an excessive number of microservices.

## Goal

To provide an easy to understand project with a non-trivial number of microservices to try out
monitoring, tracing and observability tooling.

## Technologies

- [Go](https://golang.org/)
- [gRPC](https://grpc.io/)
- [Kubernetes](https://kubernetes.io/)
- [Helm](https://helm.sh/)
- [Tilt](https://tilt.dev/)
- [TailwindCSS](https://tailwindcss.com/)
- [OpenTelemetry](https://opentelemetry.io/)
- [Lightstep](https://lightstep.com/)
- [Playwright](https://playwright.dev/)
- [Artillery](https://www.artillery.io/) with [artillery-engine-playwright](https://github.com/artilleryio/artillery-engine-playwright)

## Building and Running Locally

You can run all deployments locally using Tilt. Open a terminal in the root directory of the repo and run:

```
tilt up
```

This will build all containers and deploy a Helm chart to your current Kubernetes context, forwarding the web UI to port 8080 on localhost.

Tilt will also automatically run unit tests on any change to the source.

## Deployment

A Helm chart has been provided for direct deployment. To deploy from the command line:

```
helm repo add tic-tac-toverengineered \
    https://theothertomelliott.github.io/tic-tac-toverengineered/

helm repo update

helm install \
    --create-namespace --namespace tictactoe \
    tic-tac-toverengineered/tic-tac-toe --generate-name
```

## Secrets

Secrets are picked up by Tilt from _secrets.yaml_ (which is gitignored). Currently, only a
Lightstep access token is required.

Add the following content to _secrets.yaml_ to enable tracing:

```
lightstep:
  access_token: <access token>
```
