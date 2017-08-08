# go2readme

generate README for Go program

## Installation

```sh
$ go get github.com/gnue/go2readme/cmd/go2readme
```

## Usage

```
go2readme [OPTIONS] [dir]

Application Options:
  -w, --write     write to file
  -o, --output=   output file
  -t, --template= template file

Help Options:
  -h, --help      Show this help message

Arguments:
  dir:            directory
```

include `.go2readme/*.md`

ex.

```
.go2readme/
├── 01_Contributing.md
├── 02_Author.md
└── 03_License.md
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Run test suite with the `go test ./...` command and confirm that it passes
5. Run `gofmt -s`
6. Push to the branch (`git push origin my-new-feature`)
7. Create new Pull Request

## Author

[gnue](https://github.com/gnue)

## License

[MIT](LICENSE.txt)

