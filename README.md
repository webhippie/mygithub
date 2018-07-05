# MyGitHub

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/mygithub/status.svg)](http://github.dronehippie.de/webhippie/mygithub)
[![Stories in Ready](https://badge.waffle.io/webhippie/mygithub.svg?label=ready&title=Ready)](http://waffle.io/webhippie/mygithub)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/834ea4bc5aa24dfebebb203fd26f45f8)](https://www.codacy.com/app/webhippie/mygithub?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/mygithub&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/mygithub?status.svg)](http://godoc.org/github.com/webhippie/mygithub)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/mygithub)](http://goreportcard.com/report/github.com/webhippie/mygithub)
[![](https://images.microbadger.com/badges/image/tboerger/mygithub.svg)](http://microbadger.com/images/tboerger/mygithub "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

As I'm iterating quite often across all repositories for specific GitHub organizations I have created this small utility that currently just prints out all available repositories for specific users or organizations. Maybe I will add more features if I need it, this tool is more for personal use cases.


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.webhippie.de/misc/mygithub). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/webhippie/homebrew-webhippie).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8.

```bash
go get -d github.com/webhippie/mygithub
cd $GOPATH/src/github.com/webhippie/mygithub
make clean generate build

./mygithub -h
```


## Security

If you find a security issue please contact thomas@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
