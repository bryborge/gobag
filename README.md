# GoBag

A collection of general, yet useful, command-line tools written in Go. Some tools include:

* Web scraper
* Network scanner (Coming Soon)
* Packet sniffer (Coming Soon)
* Ping utility (Coming Soon)

## Built With

* [Cobra](https://github.com/spf13/cobra) - A library for creating powerful modern CLI applications.
* [Colly](https://github.com/gocolly/colly) - A Lightning Fast and Elegant Scraping Framework for Gophers.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

While there are a multitude of ways you can go about spinning up a development environment, we recommend using the
pre-configured docker/devcontainers environment.

*   [Docker](https://www.docker.com/) - Build, share, run, and verify applications anywhere.
*   [VS Code](https://code.visualstudio.com/) - General-purpose, extensible, code-editor.
*   [Devcontainers (VS Code Extension)](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) - Lets you use a Docker container as a full-featured development environment.
    *   More information about Devcontainers can be found [here](https://containers.dev/overview).

### Installing

1.  First, define the required environment variables by creating a `.env` file from the provided template:

    ```sh
    cp .env{.dist,}
    ```

    Update the environment variables with the values relevant to your use-case.

2.  Next, spin up your devcontainer. This can be done from VS Code's Command Palette (`Dev Containers: Rebuild`).

## Testing

TBD

## Releasing

This project uses [goreleaser](https://goreleaser.com/) for release management. The high-level release strategy is:

1.  Code up your changes
2.  Submit changes as a pull request against the `main` branch
3.  Merge pull request to `main`
4.  Tag the `main` branch with a version (follow [SemVer](https://semver.org/); e.g. `v1.2.3`)
5.  Run `goreleaser release --clean` to generate a new release with checksums and executables

In the future, this will be automated with CI/CD.
