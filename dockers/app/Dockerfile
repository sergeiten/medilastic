FROM golang:1.9 as builder

RUN mkdir -p /go/src/github.com/sergeiten/medilastic
WORKDIR /go/src/github.com/sergeiten/medilastic
COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app cmd/server/main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /go/src/github.com/sergeiten/medilastic/app .
CMD ["/app", "-host=0.0.0.0", "-port=8888"]
