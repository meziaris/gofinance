FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .
RUN ls -la
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o app ./cmd

# Creating the smallest possible Docker image for production
FROM alpine:3.18
WORKDIR /app

RUN apk update && apk --no-cache add git ca-certificates tzdata

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 10001 \
  app

# Import from builder.
COPY --from=builder --chown=app /app/app ./app

# Use an unprivileged user.
USER app:app
EXPOSE 3333
CMD ["./app"]
