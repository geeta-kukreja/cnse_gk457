FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /voter-api-starter

FROM alpine:latest AS run

WORKDIR /

COPY --from=build /voter-api-starter /voter-api-starter

EXPOSE 1080

CMD ["/voter-api-starter"]