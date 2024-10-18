# Hex to Base 64

This app is taken from this [gist](https://gist.github.com/jeffwecan/cc6b298f5363f24a21f4d73365be20f8) i found on GitHub.

The main reason for looking into this was to import [random_id](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/id#import) resources into Terraform using the base64 value.

Thing is if you have `keepers` defined you'll still have to fudge the state file :grimacing:

## Test

```shell
go test -v ./...
```

## Build

```shell
go build -C cmd/hextobase64 -o ./tmp/ .
```
