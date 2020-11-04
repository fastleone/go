FROM golang:1.15-alpine as golang
WORKDIR /go/src/app
COPY . .
RUN apk --no-cache add git
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
# Static build required so that we can safely copy the binary over.
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip
ENV TZ Europe/Moscow
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -q -r -0 /zoneinfo.zip .

FROM scratch
COPY --from=golang /go/bin/app /app
# the timezone data:
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
ENTRYPOINT ["/app"]
