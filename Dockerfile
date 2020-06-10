FROM golang:1.14.4 as builder
RUN mkdir /build
WORKDIR /build
RUN go get \
        github.com/pkg/errors \
        github.com/prometheus/client_golang/prometheus
ADD . /build/
RUN CGO_ENABLED=0 go build -o vigor-exporter .

FROM alpine:3.11.6
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/vigor-exporter /app/vigor-exporter
CMD ["/app/vigor-exporter"]
