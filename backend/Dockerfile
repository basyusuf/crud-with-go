FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY .env .
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/.env .
RUN cp /build/main .

EXPOSE 8080

CMD ["/dist/main"]