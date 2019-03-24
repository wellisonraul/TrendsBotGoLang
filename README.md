# in0984
<p>
  <a href="https://goreportcard.com/badge/github.com/gppeixoto/in0984"><img src="https://goreportcard.com/badge/github.com/gppeixoto/in0984"></img></a>
  <a href="https://golangci.com"><img src="https://golangci.com/badges/github.com/gppeixoto/in0984.svg"></img></a>
  <a href="https://godoc.org/github.com/gppeixoto/in0984"><img src="https://godoc.org/github.com/gppeixoto/in0984?status.svg" alt="GoDoc"></a>
</p>

Simple REST API that matches incoming texts with trending topics on Twitter.

## Development

### Dependencies
This project uses Go 1.9. We recommend using [update-golang](https://github.com/udhos/update-golang) to install Go.

### Quickstart
* Install [dep](https://github.com/golang/dep) for development (dependency management)
* Setup the dependencies by running `dep ensure`
* Replace the `.env.template` file with a `.env` properly configured (secrets, credentials, etc.)
* Build and serve the project locally running `make serve-local`.
