FROM golang:1.21.8-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -d -v ./...
RUN go build -o api ./cmd/app

EXPOSE 8083

CMD ["./api"]
