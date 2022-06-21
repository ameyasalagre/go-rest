
FROM golang:1.16-alpine

WORKDIR /go-rest

COPY go.mod ./
COPY go.sum ./
# RUN go mod download

COPY . .

RUN pwd

RUN go build -o /go-rest

RUN chmod +x go-rest

EXPOSE 8080

CMD [ "./go-rest" ]