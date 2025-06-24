# syntax=docker/dockerfile:1

FROM golang:tip-alpine3.22 AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o ftodaapp cmd/main.go

FROM golang:tip-alpine3.22
WORKDIR /build

COPY --from=builder /build/ftodaapp /app/ftodaapp
RUN touch /app/ftoda.db
EXPOSE 3000
CMD ["/app/ftodaapp"]