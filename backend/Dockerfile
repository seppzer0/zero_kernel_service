# build base image
FROM golang:1.20-alpine3.18 as base
ENV CGO_ENABLED=0
WORKDIR /build
RUN apk add --no-cache openjdk11-jre
COPY . .
RUN go mod download && go generate ./...

# launch the linter
FROM golangci/golangci-lint:v1.53-alpine as linter
ENV CGO_ENABLED=0
WORKDIR /src

COPY --from=base /build .
COPY .golangci.yaml .
RUN go mod download && golangci-lint run --timeout 5m

# build the binary
FROM base as build
WORKDIR /build
COPY --from=linter /src/lint_report.json .
RUN go test ./... && go build -o /build/zero-kernel-service ./cmd/backend

# build final image
FROM alpine:3.18
WORKDIR /app
COPY --from=build /build/zero-kernel-service ./zero-kernel-service
COPY config.yaml ./

CMD ["/app/zero-kernel-service"]
