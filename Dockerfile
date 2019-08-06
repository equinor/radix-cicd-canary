FROM golang:alpine3.9 as builder

RUN apk update && apk add git && apk add -y ca-certificates curl && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/equinor/radix-cicd-canary/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .
WORKDIR /go/src/github.com/equinor/radix-cicd-canary/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o ./rootfs/radix-cicd-canary
RUN adduser -D -g '' radix-cicd-canary

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/github.com/equinor/radix-cicd-canary/rootfs/radix-cicd-canary /usr/local/bin/radix-cicd-canary
USER radix-cicd-canary
ENTRYPOINT ["/usr/local/bin/radix-cicd-canary"]