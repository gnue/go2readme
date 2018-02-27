package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9b3f97c4791ee53078aea3d7017a20fd5fe9e16a = "# {{ .Name }}\n\n{{ .Description }}\n\n## Installation\n\n```sh\n$ go get {{ .ImportPath }}\n```\n\n## Usage\n\n{{if .IsCommand -}}\n```\n{{ .Usage }}\n```\n{{- else -}}\n```go\nimport \"{{ .ImportPath }}\"\n```\n{{- end}}\n{{if .Examples}}\n## Examples\n\n{{range .Examples -}}\n{{if .Name -}}\n### {{ .Name }}\n{{- end}}\n\n{{if .Play -}}\n```go\n{{ .Play }}\n```\n{{if .Output}}\nOutput:\n\n```\n{{ .Output }}\n```\n{{end}}\n{{- else -}}\n```go\n{{ .Code }}\n```\n{{end}}\n{{ .Doc }}\n{{end}}\n{{- end}}\n\n{{- range glob \".go2readme/*.md\"}}\n{{import .}}\n{{- end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"README.md"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1519695865, 1519695865515469111),
		Data:     nil,
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1518160739, 1518160739512881607),
		Data:     nil,
	}, "/assets/README.md": &assets.File{
		Path:     "/assets/README.md",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1519697061, 1519697061000000000),
		Data:     []byte(_Assets9b3f97c4791ee53078aea3d7017a20fd5fe9e16a),
	}}, "")
