FROM golang:1.20-buster as base

FROM base as dev

WORKDIR /src

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

CMD ["go", "run", "./cmd/main.go"]

FROM alpine:latest as prod

WORKDIR /app
COPY --from=dev /src/app /app/app
RUN chmod +x /app/app

CMD ["/app/app"]