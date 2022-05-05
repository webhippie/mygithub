# MyGitHub

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/mygithub?sort=semver)](https://github.com/webhippie/mygithub) [![Build Status](https://github.com/webhippie/mygithub/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/mygithub/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Docker Size](https://img.shields.io/docker/image-size/webhippie/mygithub/latest)](https://hub.docker.com/r/webhippie/mygithub) [![Docker Pulls](https://img.shields.io/docker/pulls/webhippie/mygithub)](https://hub.docker.com/r/webhippie/mygithub) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/mygithub.svg)](https://pkg.go.dev/github.com/webhippie/mygithub) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/mygithub)](https://goreportcard.com/report/github.com/webhippie/mygithub) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/834ea4bc5aa24dfebebb203fd26f45f8)](https://www.codacy.com/gh/webhippie/mygithub/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/mygithub&amp;utm_campaign=Badge_Grade)

As I'm iterating quite often across all repositories for specific GitHub
organizations I have created this small utility that currently just prints out
all available repositories for specific users or organizations. Maybe I will add
more features if I need it, this tool is more for personal use, but of course
you can use it if you like it.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/webhippie/mygithub.git
cd mygithub

make generate build

./bin/mygithub -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/webhippie/mygithub/releases
[dockerhub]: https://hub.docker.com/r/webhippie/mygithub/tags/
[quay]: https://quay.io/repository/webhippie/mygithub?tab=tags
[docs]: https://webhippie.github.io/mygithub/#getting-started
[golang]: http://golang.org/doc/install.html
