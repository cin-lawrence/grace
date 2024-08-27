FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/grace

FROM scratch

COPY --from=builder /app/grace /usr/bin/

CMD ["grace"]
