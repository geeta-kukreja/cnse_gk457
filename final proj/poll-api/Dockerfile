FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /poll-api

FROM alpine:latest AS run

WORKDIR /

COPY --from=build /poll-api /poll-api

EXPOSE 1080

CMD ["/poll-api"]
