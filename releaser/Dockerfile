FROM ubuntu:24.04
LABEL name="releaser chroot"

VOLUME /src
RUN apt update && apt upgrade -yy

RUN apt install -yy tar bash git curl zip unzip sqlite3 curl gcc make

RUN curl -SsL https://go.dev/dl/go1.22.4.linux-amd64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
RUN export PATH=$PATH:$(go env GOPATH)

WORKDIR /src

CMD ["/bin/bash", "/src/do_release.sh"]
