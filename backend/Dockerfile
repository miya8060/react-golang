FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

# airのインストール
RUN go install github.com/air-verse/air@latest
# sqlcのインストール
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

EXPOSE 8080
CMD ["air"]
