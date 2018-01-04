FROM golang:1.8-alpine as build-env
RUN apk add --no-cache git build-base && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk add --no-cache upx
ADD . /go/src/gitlab.com/ulm0/graphql-docker-api/
WORKDIR /go/src/gitlab.com/ulm0/graphql-docker-api/cmd/gdapi
RUN go get -d ./ && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo && \
#   upx --best -qq gdapi && \
    upx --ultra-brute -qq gdapi && \
    upx -t gdapi

FROM scratch
# Build-time metadata as defined at http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL maintainer="Pierre Ugaz <pierre.ugaz@ruway.me>" \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.description="A GraphQL Server for the Docker API, written in Go" \
    org.label-schema.name="GraphQL Docker API" \
    org.label-schema.schema-version="1.0" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url="https://gitlab.com/ulm0/graphql-docker-api" \
    org.label-schema.vendor="Pierre Ugaz" \
    org.label-schema.version=$VERSION
ENV UI_ENDPOINT="" \
    DOCKER_CERT_PATH="" \
    DOCKER_HOST="" \
    GQL_PORT="" \
    GRAPHIQL=""
COPY --from=build-env /go/src/gitlab.com/ulm0/graphql-docker-api/cmd/gdapi/gdapi .
EXPOSE 8080
CMD ["/gdapi"]
