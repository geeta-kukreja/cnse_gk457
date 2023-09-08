FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /votes-api

FROM alpine:latest AS run

WORKDIR /

COPY --from=build /votes-api /votes-api

EXPOSE 1082

CMD ["/votes-api"]
