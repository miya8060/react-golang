FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

# 必要なツールのインストール
RUN go install github.com/air-verse/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

EXPOSE 8080
CMD ["air"]
