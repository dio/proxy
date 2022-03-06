# proxy

See [`internal/proxy`](./internal/proxy/) for a sample usage of the public APIs
([`config`](./config/), [`handler`](./handler/), and [`runner`](./runner/)).

## Quick start

```console
go run github.com/dio/proxy/xds-server@main --resources testdata/hello
```

```console
go run github.com/dio/proxy@main --node-id testdata/hello
```

## License

See: [LICENSE](./LICENSE).
