FROM golang:latest

WORKDIR /day5-middlewares

COPY . .

RUN go build -o main .

EXPOSE 9999

CMD ["./main"]

