package web

import "embed"

// content is our static web server content.
//go:embed *
var Content embed.FS
