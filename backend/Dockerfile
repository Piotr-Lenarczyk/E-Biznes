FROM golang:1.23

RUN apt update

RUN apt install -y curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x *.sh

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
