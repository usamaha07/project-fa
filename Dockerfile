FROM golang:1.16.15

RUN mkdir /app

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o ./builder

EXPOSE 80

CMD ./builder