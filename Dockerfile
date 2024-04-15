FROM golang:1.21 as builder

WORKDIR /app
COPY . /app
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -ldflags="-s -w" -a -installsuffix cgo -o kube-pull-exporter main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/kube-pull-exporter /kube-pull-exporter
ENTRYPOINT ["/kube-pull-exporter"]
