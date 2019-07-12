FROM golang:alpine3.9 as builder

RUN apk update && apk add git && apk add -y ca-certificates curl && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/equinor/radix-cicd-canary-golang/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .
WORKDIR /go/src/github.com/equinor/radix-cicd-canary-golang/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o ./rootfs/radix-cicd-canary-golang
RUN adduser -D -g '' radix-cicd-canary-golang

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/github.com/equinor/radix-cicd-canary-golang/rootfs/radix-cicd-canary-golang /usr/local/bin/radix-cicd-canary-golang
USER radix-cicd-canary-golang
ENTRYPOINT ["/usr/local/bin/radix-cicd-canary-golang"]