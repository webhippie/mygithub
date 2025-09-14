FROM --platform=$BUILDPLATFORM golang:1.25.1-alpine3.21 AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/mygithub
COPY . /go/src/mygithub/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.22

RUN apk add --no-cache ca-certificates mailcap

ENTRYPOINT ["/usr/bin/mygithub"]
CMD ["help"]

COPY --from=builder /go/src/mygithub/bin/mygithub /usr/bin/mygithub
WORKDIR /
