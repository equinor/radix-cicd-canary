FROM docker.io/golang:1.23-alpine3.21 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /src

# Install project dependencies
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy and build project code
COPY . .
RUN go build -ldflags="-s -w" -o /build/radix-cicd-canary

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /build/radix-cicd-canary .
USER 1000
ENTRYPOINT ["/app/radix-cicd-canary"]