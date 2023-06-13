FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod file.go Text ./

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main /build/Text .

FROM scratch

COPY --from=builder /dist/main /dist/Text .

ENTRYPOINT ["/main"]