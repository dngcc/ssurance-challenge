FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o ./api ./cmd

CMD ["./api"]