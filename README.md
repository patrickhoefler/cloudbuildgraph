# cloudbuildgraph

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/cloudbuildgraph)](https://goreportcard.com/report/github.com/patrickhoefler/cloudbuildgraph)
[![Maintainability](https://api.codeclimate.com/v1/badges/e6b4c7aef80d06332d19/maintainability)](https://codeclimate.com/github/patrickhoefler/cloudbuildgraph/maintainability)

`cloudbuildgraph` visualizes your Google Cloud Build pipelines.

## Example Output

![Example graph](example/cloudbuild.png)

## Getting Started

### Prerequisites

- A `cloudbuild.yaml`, `cloudbuild.yml` or `cloudbuild.json` [Google Cloud Build config](https://cloud.google.com/cloud-build/docs/build-config) file in your current working directory

### Installation and Usage

Running `cloudbuildgraph` will create a `cloudbuild.pdf` file in your current working directory that contains a visual graph representation of your Cloud Build pipeline.

#### Docker / [nerdctl](https://github.com/containerd/nerdctl)

```shell
docker run \
  --rm \
  --workdir /workspace \
  --volume "$(pwd)":/workspace \
  ghcr.io/patrickhoefler/cloudbuildgraph
```

#### [Homebrew](https://brew.sh/)

```text
brew install patrickhoefler/tap/cloudbuildgraph
cloudbuildgraph
```

#### Build from Source

```text
go build
./cloudbuildgraph
```

## License

[MIT](https://github.com/patrickhoefler/cloudbuildgraph/blob/main/LICENSE)
