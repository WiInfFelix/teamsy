FROM golang:alpine as builder

COPY go.mod go.sum /app/

WORKDIR /app/
RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 GOOS=linux go build

FROM debian:buster

COPY --from=builder /app /app/

EXPOSE 8080/tcp

ENTRYPOINT ["/app/teamsy"]