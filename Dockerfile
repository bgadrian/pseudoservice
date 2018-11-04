FROM golang:1.11.1 AS builder
WORKDIR /src
COPY . .
RUN make build

#TODO make this work (not found error if I activate multi-stage build)
#FROM scratch
#COPY --from=builder /src/build/pseudoservice .
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["./build/pseudoservice", "--read-timeout=1s", "--write-timeout=1s", "--keep-alive=15s", "--listen-limit=1024", "--max-header-size=3KiB", "--host=0.0.0.0"]