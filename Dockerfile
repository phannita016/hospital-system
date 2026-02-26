FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY /cmd/config.yaml .

RUN go build -o server ./cmd/

EXPOSE 8080

CMD ["./server"]
