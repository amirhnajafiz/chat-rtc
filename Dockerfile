FROM golang:1.19-alpine

ARG PORT=5000

COPY . .

RUN go build -o main

EXPOSE $PORT

CMD ./main --port $PORT