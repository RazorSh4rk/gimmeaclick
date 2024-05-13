FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
VOLUME [ "/app" ]
RUN go build -o main .
EXPOSE 9001

CMD ["./main"]
