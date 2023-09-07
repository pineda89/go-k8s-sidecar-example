FROM golang as builder

WORKDIR /app

COPY . .

RUN go build -o binary

FROM debian

COPY --from=0 /app/binary .

CMD ["./binary"]