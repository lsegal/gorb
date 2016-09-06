# gorb

```sh
go<->rb
```

gorb provides a tool `gorbgen` to generate a native Ruby wrapper for a given Go
package into the `ext/<package>` directory. The module name generated will be
relative to the package directory you run `gorbgen` from.

## Installing

```sh
$ go get -a github.com/lsegal/gorb/...
$ go install github.com/lsegal/gorb/cmd/gorbgen
```

## Usage Example

Here is an example of running the built-in `Test::Fib::Fibonacci` package:

```sh
$ cd $GOPATH/github.com/lsegal/gorb
$ gorbgen -build test/fib
$ ruby ext/test/fib/test_fib.rb
```

## Run Tests

To run acceptance tests:

```sh
$ ./test/test.sh
```

Tests pass if exit code is 0.

## TODO

- [ ] Support for slices, maps, complex pointer types.
- [x] Support for pass-by-value structs (partial value type support).
- [x] Support for external packages (partial support for local packages).
- [ ] Support for error returns (as exceptions?).
- [ ] Compilation support for other systems (Windows not supported).

## License

Licensed under the BSD license. Copyright Loren Segal 2016.
