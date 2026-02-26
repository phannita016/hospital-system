FROM golang:1.24-alpine

WORKDIR /app

# Install dependencies for building and air
RUN apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

COPY . .

# Ensure /go/bin is in PATH
ENV PATH="/go/bin:${PATH}"

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
