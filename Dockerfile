FROM golang:1.25.8-alpine3.23 AS builder

ARG GOPROXY=https://goproxy.cn,direct

ENV CGO_ENABLED=0 \
    GOTOOLCHAIN=local \
    GOPROXY=${GOPROXY}

WORKDIR /src

# Cache dependencies separately from the application source.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /tmp/goweb .

FROM alpine:3.23

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S goweb \
    && adduser -S -G goweb goweb \
    && mkdir -p /app \
    && chown goweb:goweb /app

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder --chown=goweb:goweb /tmp/goweb ./goweb
COPY --chown=goweb:goweb conf ./conf
COPY --chown=goweb:goweb static ./static
COPY --chown=goweb:goweb templates ./templates

USER goweb

EXPOSE 8888

HEALTHCHECK --interval=10s --timeout=3s --start-period=10s --retries=5 \
    CMD wget -q -O /dev/null http://127.0.0.1:8888/ping || exit 1

STOPSIGNAL SIGTERM

ENTRYPOINT ["./goweb"]
