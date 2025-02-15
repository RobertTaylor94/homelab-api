FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o home_api ./cmd/homeapi/

FROM golang:alpine

WORKDIR /app

COPY --from=builder /app/home_api .

EXPOSE 8081

ENTRYPOINT [ "./home_api" ]