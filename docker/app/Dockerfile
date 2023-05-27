FROM golang:1.20.4

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

CMD ["go", "mod", "tidy"]

