### Release image
FROM ubuntu:latest@sha256:b4aa552dd3f2ed84f3214b0e8add3648aee0205ef58295c92fe0899f96ad8755

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
