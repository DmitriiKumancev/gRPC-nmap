# gRPC-nmap

![GitHub](https://img.shields.io/github/license/DmitriiKumancev/gRPC-nmap)

gRPC-nmap is a powerful gRPC service that acts as a wrapper over Nmap for network vulnerability scanning. It is designed to make it easier to perform vulnerability assessments on target hosts and retrieve detailed information about any vulnerabilities discovered.

## Features

- **gRPC Interface**: gRPC-nmap provides a simple and efficient API that allows you to scan target hosts for vulnerabilities.
- **Nmap Integration**: The service leverages Nmap and the [nmap-vulners](https://github.com/vulnersCom/nmap-vulners) script for vulnerability scanning.
- **Easy-to-Use**: With gRPC, you can easily integrate this service into your applications or use the provided client.
- **Extensible**: The service can be extended to add custom scripts or integrate with other tools for deeper analysis.

## Vulnerability Scanning

Vulnerability scanning is performed using Nmap and the `vulners` script. Below is an example command used for vulnerability scanning:

```bash
nmap -sV -p 80 --script vulners testasp.vulnweb.com
```

## Installation

To use gRPC-nmap, you need to have Nmap installed. You can download and install Nmap from the [official website](https://nmap.org/download.html).

## Usage

### Build

To create the binary file, run:

```bash
make build
```

### Run Server

To run the gRPC server, use:

```bash
make server
```

### Run Client

To run the gRPC client, use:

```bash
make client
```

### Linters

For linting your code, you can use the [golangci-lint](https://golangci-lint.run/) tool. To install it, run:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Then, run the linter using:

```bash
make lint
```

### Run Tests

To execute the tests, run:

```bash
make test
```
