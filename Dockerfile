FROM --platform=linux/amd64 golang:1.21.7-alpine as builder

WORKDIR /app
COPY . /app
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o kube-pull-exporter main.go

FROM --platform=linux/amd64 alpine:3.17

WORKDIR /app
COPY --from=builder /app/kube-pull-exporter /usr/bin/kube-pull-exporter
