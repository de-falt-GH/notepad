FROM golang:1.22 as builder

WORKDIR /app/backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o backend_app .

CMD ["./backend_app"]

