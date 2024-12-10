FROM golang:alpine AS builder

WORKDIR /proj
COPY go.mod go.mod
COPY main.go main.go

RUN go mod tidy
RUN go build -o time main.go

FROM alpine:latest

COPY --from=builder /proj/time /bin/time

ENTRYPOINT ["/bin/time"]
