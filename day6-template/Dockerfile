FROM golang:latest

ENV TZ=Asia/Shanghai

WORKDIR /day6-template

COPY . .

RUN go build -o main .

EXPOSE 9999

CMD ["./main"]

