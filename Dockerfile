FROM golang as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/main ./cmd/back-gloss/main.go

FROM alpine

WORKDIR /app
COPY --from=builder /deploy/server/ .
COPY --from=builder /app/config/ ./config/

EXPOSE 8080

ENTRYPOINT ["./main"]