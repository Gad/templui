package components

import "embed"

//go:embed **/*.templ **/*.go
var TemplFiles embed.FS
