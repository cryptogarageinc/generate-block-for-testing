FROM alpine:3.16

RUN apk add --no-cache libstdc++ wget

WORKDIR /opt/generateblock

RUN wget https://github.com/cryptogarageinc/generate-block-for-testing/releases/download/v0.0.2/generateblock-alpine3_16.gz \
  && gunzip generateblock-alpine3_16.gz \
  && mv generateblock-alpine3_16 generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD ["generateblock", "-h"]
