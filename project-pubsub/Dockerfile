FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY project-pubsub/ .
RUN go build -o project-pubsub

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/project-pubsub .
COPY .env .

# Optional: install necessary packages if required (like PostgreSQL client) 
# RUN apk add postgresql-client 

CMD ["/app/project-pubsub"]