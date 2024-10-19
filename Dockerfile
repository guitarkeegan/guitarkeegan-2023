# Use the official Go image as the base image
FROM golang:1.23.2-alpine3.19 AS build-stage

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main

# Deploy Stage
FROM alpine:3.19
RUN addgroup -S keegan && adduser -S keegan -G keegan
USER keegan
WORKDIR /app

COPY --from=build-stage --chown=keegan:keegan /build/main .
COPY --chown=keegan:keegan ./images ./images
COPY --chown=keegan:keegan ./static ./static
COPY --chown=keegan:keegan ./templates ./templates
COPY --chown=keegan:keegan ./videos ./videos
COPY --chown=keegan:keegan ./db.db .

# Set the command to run your Go application
CMD ["/app/main"]
