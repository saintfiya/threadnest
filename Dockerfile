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
RUN go build -trimpath -ldflags="-s -w" -o /tmp/threadnest .

FROM alpine:3.23

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S threadnest \
    && adduser -S -G threadnest threadnest \
    && mkdir -p /app \
    && chown threadnest:threadnest /app

ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder --chown=threadnest:threadnest /tmp/threadnest ./threadnest
COPY --chown=threadnest:threadnest conf ./conf
COPY --chown=threadnest:threadnest static ./static
COPY --chown=threadnest:threadnest templates ./templates

USER threadnest

EXPOSE 8888

HEALTHCHECK --interval=10s --timeout=3s --start-period=10s --retries=5 \
    CMD wget -q -O /dev/null http://127.0.0.1:8888/ping || exit 1

STOPSIGNAL SIGTERM

ENTRYPOINT ["./threadnest"]
