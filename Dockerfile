FROM golang:1.23-alpine AS builder


RUN apk update && apk add build-base git net-tools openssh python3 tzdata

WORKDIR /app

COPY .env.example .env
COPY . .

RUN go install github.com/buu700/gin@latest
RUN go mod tidy

RUN make build

FROM alpine:3.18

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    apk --no-cache add curl && \
    mkdir /app

WORKDIR /app

EXPOSE 8003

COPY --from=builder /app /app

ENTRYPOINT ["/app/field-service"]
