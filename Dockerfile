FROM golang:alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w"

FROM alpine
WORKDIR /app
COPY --from=builder /app/rssagg .
CMD ["/bin/sh", "-c", "./rssagg"]
