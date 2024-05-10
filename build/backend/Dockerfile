# Use a specific version of Golang for reproducibility.
FROM golang:1.22 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the rest of the backend code
COPY . .

# Build the application
RUN go build -o backend_app .

# Use a smaller base image to run the compiled binary
FROM golang:1.22-alpine
WORKDIR /app
COPY --from=builder /app/backend_app .

# Command to run the executable
CMD ["./backend_app"]
