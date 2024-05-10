FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o backend_app .

CMD ["./backend_app"]

