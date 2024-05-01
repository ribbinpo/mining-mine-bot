FROM golang:1.22.2-alpine3.19 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM debian:buster-slim as runtime

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app

COPY --from=build /app/main /bin/main

CMD ["/bin/main"]