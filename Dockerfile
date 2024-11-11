# Dockerfile to build simple_bank container image

# Build stage
FROM golang:1.21-alpine3.19 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
# Run db migration and start the server to listen on port 8080
FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/migrate ./migrate
COPY db/migration ./migration
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
EXPOSE 8080
ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
