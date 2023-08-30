# For the love of Go

## Bookstore

### Exported

types that begin with a capital letter are exported and can be used outside of the package whereas lowercase types are private to the package.

### Tests

Test coverage can be determined by running `go test -cover`

Create a coverage profile and generate an HTML output

```Shell
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

_Test behaviours not functions_ as covered is not the same as tested.