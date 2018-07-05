FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="MyGitHub" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates mailcap bash && \
  addgroup -g 1000 mygithub && \
  adduser -D -h /var/lib/mygithub -s /bin/bash -G mygithub -u 1000 mygithub

USER mygithub
ENTRYPOINT ["/usr/bin/mygithub"]
CMD ["help"]

COPY dist/binaries/mygithub-*-linux-amd64 /usr/bin/
