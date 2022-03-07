# proxy

## Installation

```console
curl https://dio.github.io/proxy/install.sh | bash -s -- -b /usr/local/bin
```

> Note: By default the script installs (if executed without `-- -b /usr/local/bin`) in the current
> directory: `./bin/proxy`.

## Running

To generate a proxy config with specified admin and stats ports:

```console
proxy --admin-port 9000 --stats-port 9001 --output a.yaml
```

Then you can run proxy with that config:

```console
proxy -- -c ./a.yaml
```

`proxy` comes with its own xDS server implementation that watches a directory with proxy configs in
it. As an example, see: [testdata/hello](./testdata/hello/).

```console
proxy --xds-resources testdata/hello --admin-port 9000
```

> Note: you can start with a non-existent directory as argument for `--xds-resources`.

This starts proxy with configurations defined in `testdata/hello/*.yaml`. Note that it watches
changes inside that directory. For example, if you change the listener address for `listener_0` in
[`testdata/hello/a.yaml`](./testdata/hello/a.yaml), e.g. from `10000` to `10001` it will immediately
reflected (this can be checked by visiting the admin interface, in this example:
https://localhost:9000)

You can test it out:

```console
curl localhost:10000
curl localhost:10001
```

If you have a running xDS server already, you can connect to that and set your proxy's `Node ID`:

```console
# localhost:9901 is the default value for `xds-server-url`.
proxy --node-id my-proxy
```

```console
proxy --node-id my-proxy --xds-server-url localhost:9902
```

> Note: TLS configuration (and other channel and call credentials) will be added in the next release.

### Using Go

`proxy` is a Go project, hence can be installed using:

```console
go install github.com/dio/proxy@latest
```

Also, you can _go run_ it directly:

```console
go run github.com/dio/proxy@latest --help
```

### Standalone xds-server

You can download standalone `xds-server` starts from [v0.0.1-rc2](https://github.com/dio/proxy/releases/tag/v0.0.1-rc2).


```console
xds-server --resources testdata/hello
```

Using Go, you can do:

```console
go run github.com/dio/proxy/xds-server@latest --help
```

## Embedding

See [`internal/proxy`](./internal/proxy/) for a sample usage of the public APIs
([`config`](./config/), [`handler`](./handler/), and [`runner`](./runner/)).


## License

See: [LICENSE](./LICENSE).
