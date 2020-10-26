# tic-tac-toverengineered

A simple Tic Tac Toe game build with an excessive number of microservices.

## Goal

To provide an easy to understand project with a non-trivial number of microservices to try out
monitoring, tracing and observability tooling.

## Technologies

* [Go](https://golang.org/)
* [Earthly](https://www.earthly.dev/)
* [gRPC](https://grpc.io/)
* [honeycomb.io](https://www.honeycomb.io/)
* [Kubernetes](https://kubernetes.io/)
* [Helm](https://helm.sh/)
* [Tilt](https://tilt.dev/)

## Building and Running Locally

You can run all deployments locally using Tilt. Open a terminal in the root directory of the repo and run:

```
tilt up
```

This will build all containers and deploy a Helm chart to your current Kubernetes context, forwarding the web UI to port 8080 on localhost.

Tilt will also automatically run unit tests on any change to the source.

Protobufs for gRPC will not automatically be rebuilt, but a resource (`protos`) is provided in Tilt that can be triggered from the Web UI.

## Deployment

Long term, the Helm chart will be published for direct deployment via Helm, but for now, Tilt
can be used to deploy direct to a cluster.

Set your Kubernetes context to your desired cluster and run `tilt up`, you will need to set
[`allow_k8s_context`](https://tilt.dev/) in the Tiltfile to permit this.

## Secrets

Secrets are picked up by Tilt from _secrets.yaml_ (which is gitignored). Currently, only a
Honeycomb API key is required.

Add the following content to _secrets.yaml_ to enable tracing:

```
honeycomb:
  api_key: <YOUR API KEY>
```