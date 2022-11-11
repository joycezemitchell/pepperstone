FROM golang:1.18

WORKDIR /app
COPY . .

RUN go build -o ./bin/pepperstonecsv .

CMD ["./bin/pepperstonecsv"]