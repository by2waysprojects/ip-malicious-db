FROM --platform=$BUILDPLATFORM golang:1.22.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG TARGETARCH TARGETOS

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o ip-malicious-db ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/ip-malicious-db /usr/local/bin/ip-malicious-db

CMD ["/usr/local/bin/ip-malicious-db"]
