FROM golang:1.13.0-alpine3.10 as builder
ENV GO111MODULE=on

ADD . /bundle-api

WORKDIR /bundle-api

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main main.go

FROM alpine:3.10

RUN apk update && apk add curl bash && rm -rf /var/cache/apk/*

COPY --from=builder /bundle-api/main /opt/opa/bundle-api

ENV PATH="/opt/opa/bundle-api:${PATH}"
# Create non privileged user, set ownership and change user
RUN addgroup opa && adduser -D -G opa opa

USER opa

WORKDIR /home/opa