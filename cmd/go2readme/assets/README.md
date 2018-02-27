# {{ .Name }}

{{ .Description }}

## Installation

```sh
$ go get {{ .ImportPath }}
```

## Usage

{{if .IsCommand -}}
```
{{ .Usage }}
```
{{- else -}}
```go
import "{{ .ImportPath }}"
```
{{- end}}
{{if .Examples}}
## Examples

{{range .Examples -}}
{{if .Name -}}
### {{ .Name }}
{{- end}}

{{if .Play -}}
```go
{{ .Play }}
```
{{if .Output}}
Output:

```
{{ .Output }}
```
{{end}}
{{- else -}}
```go
{{ .Code }}
```
{{end}}
{{ .Doc }}
{{end}}
{{- end}}

{{- range glob ".go2readme/*.md"}}
{{import .}}
{{- end}}
