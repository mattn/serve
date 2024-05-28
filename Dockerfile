# syntax=docker/dockerfile:1.4

FROM golang:1.21-alpine AS build-dev
WORKDIR /go/src/app
COPY --link go.* ./
RUN apk add --no-cache upx || \
    go version && \
    go mod download
COPY --link . .
RUN CGO_ENABLED=0 go install -buildvcs=false -trimpath -ldflags '-w -s'
RUN [ -e /usr/bin/upx ] && upx /go/bin/serve || echo
FROM scratch
COPY --link --from=build-dev /go/bin/serve /go/bin/serve
COPY --from=build-dev /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/go/bin/serve"]
