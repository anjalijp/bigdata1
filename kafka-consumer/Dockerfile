FROM --platform=linux/amd64 golang:1.23-bullseye

WORKDIR /app

COPY ./config /app/config
COPY ./models /app/models

COPY go.mod go.sum ./
RUN go mod download

COPY ./kafka-consumer .

RUN go build -o consumer . 

CMD ["sh", "-c", "/app/consumer/kafka-consumer"]