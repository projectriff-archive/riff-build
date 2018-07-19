FROM golang:1.10 as builder

WORKDIR /go/src/github.com/projectriff/riff-init

COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/
COPY ./vendor/ ./vendor/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /riff-init cmd/main.go

###########

FROM gcr.io/distroless/base:latest

# The following line forces the creation of a /tmp directory
WORKDIR /tmp

WORKDIR /

COPY --from=builder /riff-init /riff-init

ENTRYPOINT ["/riff-init"]
