FROM ubuntu:20.04@sha256:6f2fb2f9fb5582f8b587837afd6ea8f37d8d1d9e41168c90f410a6ef15fa8ce5

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
