FROM golang:latest

ENV TZ=Asia/Shanghai

WORKDIR /day5-middlewares

COPY . .

RUN go build -o main .

EXPOSE 9999

CMD ["./main"]

