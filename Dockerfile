FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o tasks-api .

EXPOSE 5001

CMD [ "./tasks-api" ]