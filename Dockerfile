FROM golang:latest

RUN go get -u github.com/beego/bee/v2

WORKDIR /graphql/src/
COPY . .

RUN go mod download

EXPOSE 3000

CMD ["bee", "run"]