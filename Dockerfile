FROM golang:1.20.1 as dev

WORKDIR /usr/app

RUN go install github.com/cosmtrek/air

COPY go.mod go.sum ./

RUN  go mod tidy

COPY . .

CMD [ "air" ]

FROM golang:1.20.1 as build

WORKDIR /usr/app

COPY . .

RUN go mod tidy
RUN mkdir bin
RUN go build ./cmd/main.go
RUN mv -f main ./bin

CMD [ "./bin/main" ]