FROM golang:latest

WORKDIR /buttons
COPY go.* ./
RUN go mod download

COPY ./ /buttons

RUN apt-get update && apt-get -y upgrade

RUN go build -o myapp ./cmd/main.go

CMD ["/buttons/myapp"]