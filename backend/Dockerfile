FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY config ./config
COPY controller ./controller
COPY docs ./docs
COPY dto ./dto
COPY middleware ./middleware
COPY model ./model
COPY repository ./repository
COPY service ./service
COPY utils ./utils
COPY main.go ./

RUN go build -o main .

FROM alpine:3.20.1

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/main .

RUN chown appuser:appgroup /app/main

USER appuser

CMD ["./main"]
