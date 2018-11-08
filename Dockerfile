FROM golang:1.11 as builder
ADD . /go/src/github.com/philproctor/confluence-cli/
RUN go get golang.org/x/net/html
RUN cd /go/src/github.com/philproctor/confluence-cli/ && go build main.go

FROM debian:jessie-slim as cert-updates
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && update-ca-certificates

FROM debian:jessie-slim
COPY --from=cert-updates /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /go/src/github.com/philproctor/confluence-cli/main /usr/local/bin/confluence-cli
ENTRYPOINT [ "confluence-cli" ]