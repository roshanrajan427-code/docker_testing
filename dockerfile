
FROM golang:1.27.1

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8082

CMD ["./server"]
