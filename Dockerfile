FROM ubuntu:latest@sha256:028d7303257c7f36c721b40099bf5004a41f666a54c0896d5f229f1c0fd99993

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
