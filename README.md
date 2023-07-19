# Overview
---
`number-crunch-2` is the evolution of the [number-crunch] application.

It splits the [number-crunch] application to two micro-services; `square-root` and `cube-root`.

You can use this application for learning Kubernetes.



# Get started
---
Clone the repo.
```shell
git clone git@github.com:cloudqubes/number-crunch-2.git
go run number-crunch.go
```

# Square-root

Return the square root of a number.

URL endpoint: `/square-root/<number>`


Run `square-root`.
```shell
cd square-root
go run main.go
```

Test `square-root`
```shell
curl http://127.0.0.1:8080/square-root/16
```

```shell
{"InputNumber":16,"SquareRoot":4}
```

# Cube-root

Return the cube root of a number.

URL endpoint: `/cube-root/<number>`

Run `cube-root`.
```shell
cd cube-root
go run main.go
```

Test `cube-root`
```shell
curl http://127.0.0.1:8080/cube-root/8
```

```shell
{"InputNumber":8,"CubeRoot":2}
```

[number-crunch]: https://github.com/cloudqubes/number-crunch