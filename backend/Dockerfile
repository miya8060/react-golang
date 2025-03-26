FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/main .

EXPOSE 8080

CMD ["/usr/local/bin/main"]