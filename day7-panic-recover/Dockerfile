FROM golang:latest

ENV TZ=Asia/Shanghai

WORKDIR /day7-panic-recover

COPY . .

RUN go build -o main .

EXPOSE 9999

CMD ["./main"]

