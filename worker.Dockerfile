FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /worker cmd/worker/main.go

CMD ["/worker"]