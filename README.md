# MyGitHub

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/mygithub/status.svg)](http://github.dronehippie.de/webhippie/mygithub)
[![Coverage Status](http://coverage.dronehippie.de/badges/webhippie/mygithub/coverage.svg)](http://coverage.dronehippie.de/webhippie/mygithub)
[![Go Doc](https://godoc.org/github.com/webhippie/mygithub?status.svg)](http://godoc.org/github.com/webhippie/mygithub)
[![Go Report](http://goreportcard.com/badge/webhippie/mygithub)](http://goreportcard.com/report/webhippie/mygithub)

As I'm iterating quite often across all repositories for specific GitHub
organizations I have created this small utility that currently just prints
out all available repositories for specific users or organizations. Maybe I
will add more features if I need it, this tool is more for personal use
cases.


## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.webhippie.de/mygithub). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie).
If you are missing an architecture just write us on our nice
[Gitter](https://gitter.im/webhippie/general) chat. Take a look at the help
output, per default we enabled an auto updater for the binary to avoid bugs
related to old versions. If you find a security issue please contact
thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/webhippie/mygithub
cd $GOPATH/src/github.com/webhippie/mygithub
make deps build

bin/mygithub -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2015-2016 Thomas Boerger <http://www.webhippie.de>
```
