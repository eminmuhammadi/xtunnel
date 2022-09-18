# xtunnel

[![](https://img.shields.io/static/v1?label=Sponsor&message=%E2%9D%A4&logo=GitHub&color=%23fe8e86)](https://github.com/sponsors/eminmuhammadi)
[![CodeQL](https://github.com/eminmuhammadi/xtunnel/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/eminmuhammadi/xtunnel/actions/workflows/codeql-analysis.yml)

xtunnel is a simple tunneling tool that allows you to create a tunnel in different ways.

Note: This tool is still in development and is not ready for production use. All connection is done over insecure channels.

## Installation

You can download binary files for each platform from [the latest releases](https://github.com/eminmuhammadi/xtunnel/releases).

## Usage

### Port Forwarding
```bash
xtunnel forward --local <local-ip>:<local-port> --remote <remote-ip>:<tcp-port> --protocol tcp
```
