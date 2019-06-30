FROM scratch
COPY cloudbuildgraph /
ENTRYPOINT ["/cloudbuildgraph"]
