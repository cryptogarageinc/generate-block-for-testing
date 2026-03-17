FROM golang:1.26-alpine3.23 as builder

WORKDIR /workspace

COPY ./ ./

RUN go mod download && go install -ldflags "-s -w" ./cmd/generateblock/ && ls -l /go/bin/

FROM alpine:3.23 as server

RUN apk add --no-cache libstdc++

WORKDIR /opt/generateblock

COPY --from=builder /go/bin/generateblock /opt/generateblock/generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD ["generateblock", "-h"]
