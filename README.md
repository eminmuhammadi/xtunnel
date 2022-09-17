# xtunnel

[![](https://img.shields.io/static/v1?label=Sponsor&message=%E2%9D%A4&logo=GitHub&color=%23fe8e86)](https://github.com/sponsors/eminmuhammadi)
[![CodeQL](https://github.com/eminmuhammadi/xtunnel/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/eminmuhammadi/xtunnel/actions/workflows/codeql-analysis.yml)

xtunnel is a simple tunneling tool that allows you to create a tunnel between your local machine and a remote server.

## Installation
You can download binary files for each platform from [the latest releases](https://github.com/eminmuhammadi/xtunnel/releases).

## Problem 
Let's say you want to create a tunnel from your local machine to a remote server. But you don't have access to the remote server, because it's behind a firewall. 

Target - The target server that you want to create a tunnel to.

Master - The master server that you want to create a tunnel from.

In this case, you can create a tunnel from your local machine to a master server, and then create a tunnel from the master server to the target server.

## Examples

### Scheme
```
(no_access:8080) -- | Firewall | -- (has_access:8080) --- | Gateway | ---- (example.com:80)
```

### Usage

1. Create a tunnel from your local machine to the master server.
 
```bash
$ xtunnel -m has_access:8080 -t example.com:80 -p tcp
```

2. Create a tunnel from the master server to the target server.

```bash
$ xtunnel -m no_access:8080 -t has_access:8080 -p tcp
```