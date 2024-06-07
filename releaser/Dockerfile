FROM alpine:3.20

LABEL name="releaser chroot"

VOLUME /src

RUN adduser ReleaseAgent -DH
RUN apk update
RUN apk add tar bash

WORKDIR /src

CMD ["/bin/bash", "/src/do_release.sh"]
