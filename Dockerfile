FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o toolbelt cmd/main.go

EXPOSE 8080
CMD ["./toolbelt"]
