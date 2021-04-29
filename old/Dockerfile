FROM golang:1.16rc1-alpine as builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .
RUN go build -o bin/app

FROM alpine:3.12

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/app /usr/local/bin/

EXPOSE 1323

CMD ["app"]
