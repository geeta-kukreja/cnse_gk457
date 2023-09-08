FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /voter-api-container

FROM alpine:latest AS run

WORKDIR /

COPY --from=build /voter-api-container /voter-api-container

EXPOSE 1081

CMD ["/voter-api-container"]
