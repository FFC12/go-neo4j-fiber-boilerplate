FROM golang:1.21.0 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

EXPOSE 3000

WORKDIR /opt/app/api
CMD ["air"]