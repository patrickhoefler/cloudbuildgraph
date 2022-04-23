### Release image
FROM ubuntu:focal-20220415@sha256:115822d64890aae5cde3c1e85ace4cc97308bb1fd884dac62f4db0a16dbddb36

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
