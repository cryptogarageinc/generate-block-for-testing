FROM alpine:3.14

RUN apk add --no-cache libstdc++ wget

WORKDIR /opt/generateblock

RUN wget https://github.com/ko-matsu/generate-block-for-testing_old/releases/download/v0.0.1-rc.0/generateblock-alpine3_14.gz \
  && gunzip generateblock-alpine3_14.gz \
  && mv generateblock-alpine3_14 generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD generateblock -h
