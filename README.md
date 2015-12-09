# readlink-proxy
## Readlink-as-a-Service

**readlink-proxy** wraps [`http.FileServer()`](https://golang.org/pkg/net/http/#FileServer) to resolve symlinks and return redirects to the canonical paths.

## Overview

**readlink-proxy** is very useful for serving versioned static content.

For example, you might provide a convenient link to the latest version of your software:

```
https://example.org/latest.tar.gz
```

You simply need to change this symlink as part of your release process. Users can check the latest version by examining the `Location` header:

```
HEAD /latest.tar.gz
```

```
HTTP/1.1 302 Found
Location: /releases/example-2.3.3.tar.gz
```

If the target is not a symlink, **readlink-proxy** serves the file normally.

## Installation

Download the [pre-built release](https://github.com/benwebber/readlink-proxy/releases) for your platform.

## Usage

**readlink-proxy** serves the contents of the current directory.

Here is an example **systemd** unit file that configures it to listen on `8000/tcp` and serve `/path/to/artifacts`.

```
[Unit]
Description=readlink-proxy

[Service]
Environment=PORT=8000
WorkingDirectory=/path/to/artifacts
ExecStart=/usr/local/bin/readlink-proxy

[Install]
WantedBy=multi-user.target
```

To bind to privileged ports (`80/tcp` or `443/tcp`), place **readlink-proxy** behind a reverse proxy like Nginx.

## Configuration

| Environment Variable | Default | Description                        |
|----------------------|---------|------------------------------------|
| `PORT`               | `3000`  | port **readlink-proxy** listens on |

## License

MIT
