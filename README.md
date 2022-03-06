# proxy

See [`internal/proxy`](./internal/proxy/) for a sample usage of the public APIs
([`config`](./config/), [`handler`](./handler/), and [`runner`](./runner/)).

## Quick start

Run the following in two separate terminal sessions:

```console
go run github.com/dio/proxy/xds-server@main --resources testdata/hello
```

```console
go run github.com/dio/proxy@main --node-id testdata/hello
```

Check if the server is ready:

```
curl localhost:10000
curl localhost:10001
```

See: [`testdata/hello`](./testdata/hello/) for available configurations for this quick start section.

## License

See: [LICENSE](./LICENSE).
