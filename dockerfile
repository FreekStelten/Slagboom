FROM golang:latest

WORKDIR /slagboom

COPY . .

RUN go build -o main .

CMD ["./main"]
