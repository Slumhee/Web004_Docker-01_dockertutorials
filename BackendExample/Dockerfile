FROM golang:1.23-alpine

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080
# GO GO GO
CMD ["./main"]

