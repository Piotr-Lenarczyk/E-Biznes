FROM golang:1.23

RUN apt update && apt install -y curl && apt clean

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x ./*.sh && go build -o main .

EXPOSE 8080

CMD ["./main"]
