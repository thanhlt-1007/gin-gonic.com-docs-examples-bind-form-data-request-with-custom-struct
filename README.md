# gin-gonic.com-docs-examples-bind-form-data-request-with-custom-struct

- Bind form-data request with custom struct

- Reference: https://gin-gonic.com/docs/examples/bind-form-data-request-with-custom-struct/

## gvm

```sh
gvm install go1.23.5
gvm use go1.23.5
```

## go mod

```sh
go mod download
```

## go run

```sh
go run .
```

## curl

- root

```sh
curl --location 'localhost:8080'
```

- Get A

```sh
curl --location 'localhost:8080/get-a?field_a=A&field_b=B'
```

- Get B

```sh
curl --location 'localhost:8080/get-b?field_a=A&field_b=B'
```

- Get C

```sh
curl --location 'localhost:8080/get-c?field_a=A&field_c=C'
```
