# Overview

`number-crunch-2` is the evolution of the [number-crunch] application.

It splits the [number-crunch] application into two micro-services; `square-root` and `cube-root`.

`number-crunch-2` serves Kubernetes learning. Read these blog posts for what you can do with it.

1. [How to use Kubernetes ingress for routing HTTP traffic][kubernetes-ingress].
2. [Canary deployments with Ingress Nginx controller](https://cloudqubes.com/blog/canary-deployments-with-ingress-nginx/)

# Running the applications locally

Clone the repo and run the application.
```shell
git clone git@github.com:cloudqubes/number-crunch-2.git
go run number-crunch.go
```

# Square-root

The `square-root` micro-service returns the square root of a number.

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

HTTP Response
```shell
{"InputNumber":16,"SquareRoot":4}
```

## Log

`square-root` logs every http request.

Request
```shell
curl localhost:8080/square-root/4
```

Log
```shell
[<hostname>:<timestamp>:square-root v2] Request: /square-root/4, Response: {"InputNumber":4,"SquareRoot":2}
```



# Cube-root

The `cube-root` micro-service returns the cube root of a number.

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
[kubernetes-ingress]: https://cloudqubes.com/blog/how-to-use-ingress/