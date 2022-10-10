FROM golang:1.19

WORKDIR /app

ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ADD . .

RUN go build -mod vendor -o worker cmd/worker/main.go

EXPOSE 3001

CMD [ "./worker" ]