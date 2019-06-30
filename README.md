# Cloud Build Graph

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/cloudbuildgraph)](https://goreportcard.com/report/github.com/patrickhoefler/cloudbuildgraph)

Visualize your Google Cloud Build steps with GraphViz.

## Dependencies

- [GraphViz](https://www.graphviz.org/)

## Usage

### Local

```shell
cat cloudbuild.yaml | cloudbuildgraph | dot -Tpdf > cloudbuild.pdf
open cloudbuild.pdf
```

### Docker

```shell
cat cloudbuild.yaml | docker run --rm -i patrickhoefler/cloudbuildgraph | dot -Tpdf > cloudbuild.pdf
open cloudbuild.pdf
```
