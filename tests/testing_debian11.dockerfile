FROM debian:11-slim

RUN apt-get update && apt-get install -y wget

WORKDIR /opt/generateblock

RUN wget https://github.com/ko-matsu/generate-block-for-testing_old/releases/download/v0.0.1-rc.0/generateblock-linux_amd64.gz \
  && gunzip generateblock-linux_amd64.gz \
  && mv generateblock-linux_amd64 generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD generateblock -h
