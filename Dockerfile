FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

# Change PORT
EXPOSE 8000

ENTRYPOINT ["/app/binary"]