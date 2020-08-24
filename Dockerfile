FROM alpine:3.10.5@sha256:f0e9534a598e501320957059cb2a23774b4d4072e37c7b2cf7e95b241f019e35

RUN \
  apk add --update --no-cache \
  graphviz \
  ttf-freefont

COPY cloudbuildgraph /

WORKDIR /cloudbuild

ENTRYPOINT ["/cloudbuildgraph"]
