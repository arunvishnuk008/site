#build stage
FROM golang:1.27.4-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o prerise cmd/main.go

#Run stage

FROM alpine:latest

WORKDIR /app

#copy from previous step
COPY --from=builder /app/prerise .
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip

EXPOSE 8082
CMD ["./prerise"]
