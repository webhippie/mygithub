---
title: "Getting Started"
date: 2022-05-04T00:00:00+00:00
anchor: "getting-started"
weight: 20
---

## Installation

So far we are offering only a few different variants for the installation. You
can choose between [Docker][docker] or pre-built binaries which are stored on
our download mirror and GitHub releases. Maybe we will also provide system
packages for the major distributions later if we see the need for it.

### Docker

Generally we are offering the images through
[quay.io/webhippie/mygithub][quay] and [webhippie/mygithub][dockerhub], so
feel free to choose one of the providers. Maybe we will come up with Kustomize
manifests or some Helm chart.

### Binaries

Simply download a binary matching your operating system and your architecture
from our [downloads][downloads] or the GitHub releases and place it within your
path like `/usr/local/bin` if you are using macOS or Linux.

## Configuration

We provide overall three different variants of configuration. The variant based
on environment variables and commandline flags are split up into global values
and command-specific values.

### Envrionment variables

If you prefer to configure the service with environment variables you can see
the available variables below.

#### Global

MYGITHUB_LOG_LEVEL
: Set logging level, defaults to `info`

MYGITHUB_LOG_COLOR
: Enable colored logging, defaults to `true`

MYGITHUB_LOG_PRETTY
: Enable pretty logging, defaults to `true`

MYGITHUB_TOKEN
: Token to access the GitHub API

#### List

MYGITHUB_LIST_FORMAT
: Output format, could be emtpy, `name` or `json`

### Commandline flags

If you prefer to configure the service with commandline flags you can see the
available variables below.

#### Global

--log-level
: Set logging level, defaults to `info`

--log-color
: Enable colored logging, defaults to `true`

--log-pretty
: Enable pretty logging, defaults to `true`

--token
: Token to access the GitHub API

#### List

--format
: Output format, could be emtpy, `name` or `json`

### Configuration file

So far we support multiple file formats like `json` or `yaml`, if you want to
get a full example configuration just take a look at [our repository][repo],
there you can always see the latest configuration format. These example configs
include all available options and the default values. The configuration file
will be automatically loaded if it's placed at
`/etc/mygithub/config.yml`, `${HOME}/.mygithub/config.yml` or
`$(pwd)/mygithub/config.yml`.

## Usage

The program provides a few sub-commands on execution. The available config
methods have already been mentioned above. Generally you can always see a
formated help output if you execute the binary similar to something like
 `mygithub --help`.

[docker]: https://www.docker.com/
[quay]: https://quay.io/repository/webhippie/mygithub
[dockerhub]: https://hub.docker.com/r/webhippie/mygithub
[downloads]: https://dl.webhippie.de/
[repo]: https://github.com/webhippie/mygithub/tree/master/config
