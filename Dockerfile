FROM golang:1.21-alpine as base

FROM base as dev

EXPOSE 8000

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . ./

RUN go mod download

RUN ./swagger.sh

CMD ["air"]