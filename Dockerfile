FROM golang:1.11.1 AS builder
WORKDIR /src
COPY . .
RUN make build

#TODO make this work (not found error if I activate multi-stage build)
#FROM scratch
#COPY --from=builder /src/build/pseudoservice .
ENTRYPOINT ["./build/pseudoservice"]