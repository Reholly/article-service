FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gettext

COPY go.* ./
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/article-service cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /app/bin/article-service /
COPY --from=builder /app/internal/migration /migration
COPY --from=builder /app/config.yaml config.yaml

ENTRYPOINT ["./article-service"]