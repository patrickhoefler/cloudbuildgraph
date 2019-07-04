FROM scratch
COPY cloudbuildgraph /
WORKDIR /cloudbuild
ENTRYPOINT ["/cloudbuildgraph"]
