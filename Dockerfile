FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D mygoappuser1

WORKDIR /app

COPY --from=builder /app/main /app/main

ENV PORT=8083
EXPOSE ${PORT}

USER mygoappuser1

CMD ["/app/main"]

