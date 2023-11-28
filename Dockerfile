# Use the official Go image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the rest of your application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the command to run your Go application
CMD ["./main"]