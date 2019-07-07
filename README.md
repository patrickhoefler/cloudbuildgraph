# cloudbuildgraph

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/cloudbuildgraph)](https://goreportcard.com/report/github.com/patrickhoefler/cloudbuildgraph)
[![Maintainability](https://api.codeclimate.com/v1/badges/e6b4c7aef80d06332d19/maintainability)](https://codeclimate.com/github/patrickhoefler/cloudbuildgraph/maintainability)

Visualize your Google Cloud Build steps with GraphViz.

![Example graph](example/cloudbuild.png)

## Getting Started

### Prerequisites

Make sure you have [GraphViz](https://www.graphviz.org/) installed.

For now, you need a `cloudbuild.yaml` or `cloudbuild.yml` file in the working directory. Cloud Build config files in JSON are not supported yet.

### Usage

```shell
docker run --rm -v "$(pwd)":/cloudbuild patrickhoefler/cloudbuildgraph | dot -Tpdf > cloudbuild.pdf
```
