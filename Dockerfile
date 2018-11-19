FROM golang:1.11.1 AS builder
WORKDIR /src

#avoid downloading the dependencies on succesive builds
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM alpine
COPY --from=builder /src/build/pseudoservice .
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["./pseudoservice", "--read-timeout=1s", "--write-timeout=1s", "--keep-alive=15s", "--listen-limit=1024", "--max-header-size=3KiB", "--host=0.0.0.0"]