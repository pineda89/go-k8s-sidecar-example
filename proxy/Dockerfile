FROM golang as builder

WORKDIR /app

COPY . .

RUN go build -o binary

FROM debian

RUN apt-get -y update; apt-get -y install curl

COPY --from=0 /app/binary .

CMD ["./binary"]