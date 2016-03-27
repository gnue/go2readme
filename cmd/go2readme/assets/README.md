# {{ .Name }}

{{ .Description }}

## Installation

```sh
$ go get {{ .ImportPath }}
```

## Usage
{{if .IsCommand }}
```
{{ .Usage }}
```
{{else}}
```go
import "{{ .ImportPath }}"
```
{{end}}
{{if .Examples}}
## Examples
{{range .Examples}}
{{ .Doc }}{{if .Play }}
```go
{{ .Play }}
```
{{if .Output }}
```
{{ .Output }}
```
{{end}}
{{else}}
```go
{{ .Code }}
```
{{end}}{{end}}{{end}}
{{range glob ".go2readme/*.md"}}{{import .}}
{{end}}
