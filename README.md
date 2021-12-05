<p align="center"><img src="docs/sources/logo_and_name.png" alt="Loki Logo"></p>

<a href="https://drone.grafana.net/grafana/loki"><img src="https://drone.grafana.net/api/badges/grafana/loki/status.svg" alt="Drone CI" /></a>
<a href="https://circleci.com/gh/grafana/loki/tree/master"><img src="https://circleci.com/gh/grafana/loki.svg?style=shield&circle-token=618193e5787b2951c1ea3352ad5f254f4f52313d" alt="CircleCI" /></a>
<a href="https://goreportcard.com/report/github.com/grafana/loki"><img src="https://goreportcard.com/badge/github.com/grafana/loki" alt="Go Report Card" /></a>
<a href="https://slack.grafana.com/"><img src="https://img.shields.io/badge/join%20slack-%23loki-brightgreen.svg" alt="Slack" /></a>

# Loki: like Prometheus, but for logs.

> Fork of Loki 2.2.1 (apache2 licensed)

Loki is a horizontally-scalable, highly-available, multi-tenant log aggregation system inspired by [Prometheus](https://prometheus.io/).
It is designed to be very cost effective and easy to operate.
It does not index the contents of the logs, but rather a set of labels for each log stream.

Compared to other log aggregation systems, Loki:

- does not do full text indexing on logs. By storing compressed, unstructured logs and only indexing metadata, Loki is simpler to operate and cheaper to run.
- indexes and groups log streams using the same labels you’re already using with Prometheus, enabling you to seamlessly switch between metrics and logs using the same labels that you’re already using with Prometheus.
- is an especially good fit for storing [Kubernetes](https://kubernetes.io/) Pod logs. Metadata such as Pod labels is automatically scraped and indexed.
- has native support in Grafana (needs Grafana v6.0).

A Loki-based logging stack consists of 3 components:

- `promtail` is the agent, responsible for gathering logs and sending them to Loki.
- `loki` is the main server, responsible for storing logs and processing queries.
- [Grafana](https://github.com/grafana/grafana) for querying and displaying the logs.

Loki is like Prometheus, but for logs: we prefer a multidimensional label-based approach to indexing, and want a single-binary, easy to operate system with no dependencies.
Loki differs from Prometheus by focusing on logs instead of metrics, and delivering logs via push, instead of pull.

## Getting started

* [Installing Loki](https://grafana.com/docs/loki/latest/installation/)
* [Installing Promtail](https://grafana.com/docs/loki/latest/clients/promtail/installation/)
* [Getting Started Guide](https://grafana.com/docs/loki/latest/getting-started/)

## Upgrading

* [Upgrading Loki](https://grafana.com/docs/loki/latest/operations/upgrade/)

### Documentation

* [v2.2.1](https://github.com/grafana/loki/tree/v2.2.1/docs/README.md)

### Building from source

Loki can be run in a single host, no-dependencies mode using the following commands.

You need `go`, we recommend using the version found in [our build Dockerfile](https://github.com/grafana/loki/blob/master/loki-build-image/Dockerfile)

```bash

$ go get github.com/metrico/loki-apache
$ cd $GOPATH/src/github.com/metrico/loki-apache # GOPATH is $HOME/go by default.

$ go build ./cmd/loki
$ ./loki -config.file=./cmd/loki/loki-local-config.yaml
...
```

To build Promtail on non-Linux platforms, use the following command:

```bash
$ go build ./cmd/promtail
```

On Linux, Promtail requires the systemd headers to be installed for
Journal support.

With Journal support on Ubuntu, run with the following commands:

```bash
$ sudo apt install -y libsystemd-dev
$ go build ./cmd/promtail
```

With Journal support on CentOS, run with the following commands:

```bash
$ sudo yum install -y systemd-devel
$ go build ./cmd/promtail
```

Otherwise, to build Promtail without Journal support, run `go build`
with CGO disabled:

```bash
$ CGO_ENABLED=0 go build ./cmd/promtail
```

## License

Apache License 2.0, see [LICENSE](LICENSE).
