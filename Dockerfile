### Release image
FROM ubuntu:noble-20250127@sha256:72297848456d5d37d1262630108ab308d3e9ec7ed1c3286a32fe09856619a782

LABEL org.opencontainers.image.source="https://github.com/patrickhoefler/cloudbuildgraph"

RUN \
  # Install Graphviz
  # Note that the lack of a "lock" mechanism for apt dependencies
  # currently prevents a fully reproducible build
  apt-get update \
  && apt-get install -y --no-install-recommends \
  graphviz \
  && rm -rf /var/lib/apt/lists/* \
  \
  # Add a non-root user
  && useradd app

# Run as non-root user
USER app

# This currently only works with goreleaser
# or if you manually copy the binary into the main project directory
COPY cloudbuildgraph /

ENTRYPOINT ["/cloudbuildgraph"]
