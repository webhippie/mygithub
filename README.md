# MyGitHub

[![General Workflow](https://github.com/webhippie/mygithub/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/mygithub/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/834ea4bc5aa24dfebebb203fd26f45f8)](https://app.codacy.com/gh/webhippie/mygithub/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/mygithub.svg)](https://pkg.go.dev/github.com/webhippie/mygithub) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/mygithub)](https://goreportcard.com/report/github.com/webhippie/mygithub) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/webhippie/mygithub) [![Hosted By: Cloudsmith](https://img.shields.io/badge/OSS%20hosting%20by-cloudsmith-blue?logo=cloudsmith&style=flat-square)](https://cloudsmith.com)

As I'm iterating quite often across all repositories for specific GitHub
organizations I have created this small utility that currently just prints out
all available repositories for specific users or organizations. Maybe I will add
more features if I need it, this tool is more for personal use, but of course
you can use it if you like it.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. Besides that we also prepared repositories for
DEB and RPM packages which can be found at [Cloudsmith][pkgrepo]. If you prefer
to use containers you could use our images published on [GHCR][ghcr],
[Docker Hub][dockerhub] or [Quay][quay]. If you need further guidance how to
install this take a look at our [documentation][docs].

Package repository hosting is graciously provided by [Cloudsmith][cloudsmith].
Cloudsmith is the only fully hosted, cloud-native, universal package management
solution, that enables your organization to create, store and share packages in
any format, to any place, with total confidence.

## Prerequisites

We use [mise][mise] to manage all required tools and their versions. Install it
by following the [official installation instructions][mise-install], then run
the following commands inside the repository to activate mise and install all
tools defined in `mise.toml`:

```console
mise trust
mise install
```

## Build

Since all required commands ar part of our [go-task][gotask] taskfile the
commands you got to execute are quite simple:

```console
git clone https://github.com/webhippie/mygithub.git
cd mygithub

task build
./bin/mygithub -h
```

## Development

To start developing on this project you have to execute only a few commands in
multiple terminal tabs or windows:

```console
task watch
```

After that you can simply execute the tool via `bin/mygithub -h`. Generally it
supports hot reloading which means the binary gets automatically recompiled on
code changes.

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Generally we are following [conventional commits][commits] when we apply
changes. That way we are able to generate proper changelogs for every release.
Please use always pull requests to integrate new functionalities or to fix
issues.

For the release process we are following [semantic versioning][semver] which
clearly indicates if a new version just resolves bugs, includes new features or
even includes breaking changes.

After installing the tools via `mise install` as described above set up the
pre-commit hooks so they run automatically on every commit:

```console
prek install --hook-type pre-commit --hook-type commit-msg
```

> `prek` is managed by mise and will be available after `mise install`.

If you have changed something on the source you should simply commit following
the mentioned conventions:

```console
git checkout -b feat/new-feature
git add --all
git commit -m 'feat: added awesome new feature'
git push --set-upstream origin feat/new-feature
```

After pushing your changes into the Git repository you should create a pull
request on GitHub. If the pull request have been merged and everything built
fine it will also create automatically a new release at least once a week.

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[pkgrepo]: https://cloudsmith.io/~webhippie/repos/general/groups/
[releases]: https://github.com/webhippie/mygithub/releases
[downloads]: https://dl.webhippie.de/#mygithub/
[ghcr]: https://github.com/webhippie/mygithub/pkgs/container/mygithub
[dockerhub]: https://hub.docker.com/r/webhippie/mygithub/tags/
[quay]: https://quay.io/repository/webhippie/mygithub?tab=tags
[docs]: https://webhippie.github.io/mygithub/#getting-started
[cloudsmith]: https://cloudsmith.com/
[gotask]: https://taskfile.dev/installation/
[devcontainer]: https://containers.dev/
[mise]: https://mise.jdx.dev/
[mise-install]: https://mise.jdx.dev/getting-started.html
[commits]: https://www.conventionalcommits.org/en/v1.0.0/
[semver]: https://semver.org/
