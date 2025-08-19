FROM debian:13-slim

RUN apt-get update && apt-get install -y wget

WORKDIR /opt/generateblock

RUN wget --no-check-certificate https://github.com/cryptogarageinc/generate-block-for-testing/releases/download/v0.0.7/generateblock-linux_amd64.gz \
  && gunzip generateblock-linux_amd64.gz \
  && mv generateblock-linux_amd64 generateblock

RUN chmod +x /opt/generateblock/generateblock

ENV PATH $PATH:/opt/generateblock

CMD ["generateblock", "-h"]
