FROM golang:latest

LABEL maintainer="Egor Kostetskiy <kosegor@gmail.com>"

COPY . /go/src/github.com/egorkos/minesweeper

WORKDIR /go/src/github.com/egorkos/minesweeper

RUN go get ./...
RUN go build -o minesweeper cmd/minesweeper/main.go

EXPOSE 8080

CMD ["./minesweeper"]

