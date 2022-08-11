FROM golang:alpine as builder

WORKDIR /build

COPY . .
RUN go build

FROM alpine

WORKDIR /app
COPY --from=builder /build/events /app/events

CMD ["./events"]
