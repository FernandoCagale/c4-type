# c4-type [Flow](https://github.com/FernandoCagale/c4-kustomize)

### Docker

`running docker multi-stage builds and publish c4-type to HTTP and gRPC`

```sh
$   ./scripts/publish-grpc.sh
```

```sh
$   ./scripts/publish-http.sh
```

### Kubernetes and Istio - [YAML](https://github.com/FernandoCagale/c4-kustomize/tree/master/c4-type/base)

    *   service-grpc.yaml
    *   deployment-grpc.yaml
    *   service-http.yaml
    *   deployment-http.yaml

# Running local

### Standard Go Project [Layout](https://github.com/golang-standards/project-layout)

```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download "dependency injection"` [wire](https://github.com/google/wire)

```sh
$   go get -u github.com/google/wire/cmd/wire
```

```sh
$   ./scripts/start-grpc.sh
```

```sh
$   ./scripts/start-http.sh
```