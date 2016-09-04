# gorb

```
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

# Known Issues / TODO

- [ ] Support for slices, maps, complex pointer types.
- [ ] Support for external packages.
- [ ] Compilation support for other systems (Windows not supported).

# License

Licensed under the BSD license. Copyright Loren Segal 2016.
