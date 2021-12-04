FROM golang:1.17.3-buster AS builder

WORKDIR "/app"

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ../internal internal
COPY ../cmd cmd
COPY ../Makefile .

ARG COMMIT

RUN make build

FROM debian:buster

USER nobody

COPY --from=builder /app/server /app/server

EXPOSE 5000

WORKDIR /app

ENTRYPOINT ["/app/server"]

