### Release image
FROM ubuntu:focal-20220531@sha256:397930ce1f68f54e10df2fede8e268fe086e2a7e58d83e532fb3abacfa040da7

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
