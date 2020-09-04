FROM ubuntu:latest@sha256:31dfb10d52ce76c5ca0aa19d10b3e6424b830729e32a89a7c6eee2cda2be67a5

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
  graphviz \
  && rm -rf /var/lib/apt/lists/*

# Run as non-root user
RUN groupadd --gid 1000 cloudbuildgraph \
  && useradd --uid 1000 --gid cloudbuildgraph --shell /bin/bash --create-home cloudbuildgraph
USER cloudbuildgraph

# This currently only works with goreleaser
COPY cloudbuildgraph /

WORKDIR /cloudbuild

ENTRYPOINT ["/cloudbuildgraph"]
