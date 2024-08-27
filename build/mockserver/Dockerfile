FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/mockserver

FROM scratch

COPY --from=builder /app/mockserver /usr/bin/

CMD ["mockserver"]
