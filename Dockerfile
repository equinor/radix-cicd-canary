FROM --platform=$BUILDPLATFORM docker.io/golang:1.24.2-alpine3.21 AS builder
ARG TARGETARCH
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

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