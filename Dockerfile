FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /api-gateway

EXPOSE 8081

ENTRYPOINT [ "/api-gateway" ]