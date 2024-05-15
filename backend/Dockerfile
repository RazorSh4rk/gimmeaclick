# FROM golang:latest

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main .
# EXPOSE 9001

# CMD ["./main"]


FROM golang:latest as builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
WORKDIR /app
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o gcp-data-drive


FROM gcr.io/distroless/static
COPY --from=builder /app/gcp-data-drive /gcp-data-drive

EXPOSE 9001

CMD ["/gcp-data-drive"]