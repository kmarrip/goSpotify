FROM golang:1.16-alpine

WORKDIR /app

ADD . ./

RUN go mod download

RUN go build -o ./app

CMD ./app