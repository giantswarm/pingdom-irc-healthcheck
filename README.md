# pingdom-irc-healthcheck

[![](https://img.shields.io/docker/pulls/giantswarm/pingdom-irc-healthcheck.svg)](http://hub.docker.com/r/giantswarm/pingdom-irc-healthcheck) [![IRC Channel](https://img.shields.io/badge/irc-%23giantswarm-blue.svg)](https://kiwiirc.com/client/irc.freenode.net/#giantswarm)


`pingdom-irc-healthcheck` is a microservice, written in Go, that provides a Pingdom HTTP Custom check endpoint which can be used to test if a user is in IRC.

Works well with [slack-ircbridge](https://github.com/giantswarm/slack-ircbridge).

## Getting Project

Clone the git repository: https://github.com/giantswarm/pingdom-irc-healthcheck.git

Download the latest docker image from here: https://hub.docker.com/r/giantswarm/pingdom-irc-healthcheck/

### How to build

#### Dependencies

- [github.com/juju/errgo](https://github.com/juju/errgo)
- [github.com/thoj/go-ircevent](https://github.com/thoj/go-ircevent)

#### Building the standard way

- Use `make` to build the binary.
- To build the Docker image, build the binary for linux using `GOOS=linux make`, and then build the image with `make docker`.

## Running PROJECT

- `pingdom-irc-healthcheck` is currently designed to work in a Giant Swarm cluster. A Docker image is also available, for running in other environments.

## Contact

- Mailing list: [giantswarm](https://groups.google.com/forum/!forum/giantswarm)
- IRC: #[giantswarm](irc://irc.freenode.org:6667/#giantswarm) on freenode.org
- Bugs: [issues](https://github.com/giantswarm/pingdom-irc-healthcheck/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches, the contribution workflow as well as reporting bugs.

## License

PROJECT is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.