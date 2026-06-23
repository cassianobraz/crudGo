FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o meuprimeirocrudgo .

FROM alpine:latest AS runner

WORKDIR /app

RUN adduser -D huncoding

COPY --from=builder /app/meuprimeirocrudgo /app/meuprimeirocrudgo

RUN chown -R huncoding:huncoding /app
RUN chmod +X /app/meuprimeirocrudgo

EXPOSE 8080

USER huncoding

CMD ["/app/meuprimeirocrudgo"]