FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
ADD .env .
COPY . .
RUN go build -o main src/main.go


FROM alpine

WORKDIR /build
COPY --from=builder /build/main /build/main

CMD ["./main"]