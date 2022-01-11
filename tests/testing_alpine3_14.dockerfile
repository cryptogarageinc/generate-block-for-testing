FROM alpine:3.14

RUN apk add --no-cache libstdc++ wget

WORKDIR /opt/generateblock

RUN wget https://github.com/cryptogarageinc/generate-block-for-testing/releases/download/v0.0.1-rc.3/generateblock-alpine3_14.gz \
  && gunzip generateblock-alpine3_14.gz \
  && mv generateblock-alpine3_14 generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD ["generateblock", "-h"]
