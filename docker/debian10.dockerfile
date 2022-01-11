FROM golang:1.16-buster as builder

WORKDIR /workspace

COPY ./ ./

RUN go mod download && go install -ldflags "-s -w" ./cmd/generateblock/ && ls -l /go/bin/

FROM debian:10-slim as server

WORKDIR /opt/generateblock

COPY --from=builder /go/bin/generateblock /opt/generateblock/generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD ["generateblock", "-h"]
