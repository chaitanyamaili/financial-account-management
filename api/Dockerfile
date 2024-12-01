# Use the official Golang image for building the app
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory to the working directory
COPY . .

# Build the Go application
RUN go build -o main main.go

# Expose the application's port
EXPOSE 8080

# Run the application
ENTRYPOINT [ "/app/main" ]