# syntax=docker/dockerfile:1

FROM golang:latest AS builder
ENV GOOS=linux

WORKDIR /build

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      build-essential \
      libsqlite3-dev

ADD go.mod .
COPY . .
RUN go build -o app cmd/main.go

FROM golang:latest
WORKDIR /build

COPY --from=builder /build/app /build/ftoda
COPY --from=builder /build/templates /build/templates
EXPOSE 3000
CMD ["/build/ftoda"]