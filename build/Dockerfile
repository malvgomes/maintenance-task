FROM golang:1.19

WORKDIR /app

ADD . .

RUN go build -mod vendor -o maintenance-task cmd/main.go

EXPOSE 3000

CMD [ "./maintenance-task" ]