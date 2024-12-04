FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main cmd/http/main.go

EXPOSE 3000

CMD [ "./main" ]