FROM golang:latest

WORKDIR /app

COPY ./ ./
COPY ./cmd/.env ./

RUN go build -o main ./cmd
CMD [ "./main" ]