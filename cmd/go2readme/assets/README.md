# {{ .Name }}

{{ .Description }}

## Installation

{{if .IsCommand -}}
```sh
$ go install {{ .ImportPath }}@latest
```
{{- else -}}
```sh
$ go get {{ .ImportPath }}
```
{{- end}}

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
