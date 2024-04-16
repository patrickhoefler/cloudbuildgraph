### Release image
FROM ubuntu:focal-20240410@sha256:f71981ca44fed32f9c86886c5abc6db23363e2f8e6a99e4d03687444939427d0

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
