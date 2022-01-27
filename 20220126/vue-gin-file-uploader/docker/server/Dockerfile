FROM golang:1.11

WORKDIR /api
COPY ./server .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
