FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV GRPC_SRV_ADDR :50052
ENV DB_USER root
ENV DB_PASS admin

RUN go build

EXPOSE 50052

CMD ["./grpc-demo"]