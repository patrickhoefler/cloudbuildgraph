FROM ubuntu:latest@sha256:1d7b639619bdca2d008eca2d5293e3c43ff84cbee597ff76de3b7a7de3e84956

LABEL org.opencontainers.image.source="https://github.com/patrickhoefler/cloudbuildgraph"

# Because there is no "lock" mechanism for apt dependencies,
# this step prevents a fully reproducible build
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
  graphviz \
  && rm -rf /var/lib/apt/lists/*

# Run as non-root user
RUN groupadd --gid 1000 cloudbuildgraph \
  && useradd --uid 1000 --gid cloudbuildgraph --shell /bin/bash --create-home cloudbuildgraph
USER cloudbuildgraph

# This currently only works with goreleaser
# or if you manually copy the binary into the main project directory
COPY cloudbuildgraph /

WORKDIR /cloudbuild

ENTRYPOINT ["/cloudbuildgraph"]
