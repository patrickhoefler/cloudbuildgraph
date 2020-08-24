FROM alpine:3.12.0@sha256:a15790640a6690aa1730c38cf0a440e2aa44aaca9b0e8931a9f2b0d7cc90fd65

RUN \
  apk add --update --no-cache \
  graphviz \
  ttf-freefont

COPY cloudbuildgraph /

WORKDIR /cloudbuild

ENTRYPOINT ["/cloudbuildgraph"]
