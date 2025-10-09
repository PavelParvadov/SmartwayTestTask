
FROM golang:1.25.0-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o employee_service cmd/employee-service/main.go && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migrator cmd/migrator/main.go


FROM alpine:3.20
WORKDIR /app

RUN adduser -D -H appuser && chown -R appuser /app
USER appuser

COPY --from=builder /app/employee_service ./
COPY --from=builder /app/migrator ./
COPY ./migrations ./migrations

ENV PORT=8081

CMD ["./employee_service"]



