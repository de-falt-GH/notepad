# Use a specific version of Golang for reproducibility.
FROM golang:1.21 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the rest of the backend code
COPY . .

# Build the application
RUN go build -o main .

# Use a smaller base image to run the compiled binary
FROM golang:1.21-alpine
WORKDIR /app
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
