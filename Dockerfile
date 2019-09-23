# builder
FROM golang:1.12 as builder
WORKDIR /go/src/github.com/jasonmccallister/ip
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o ip
# runtime
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/jasonmccallister/ip/ip /ip
CMD ["/ip"]
