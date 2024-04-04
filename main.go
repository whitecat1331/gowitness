package gowitness

import (
	"embed"

	"github.com/whitecat1331/gowitness/cmd"
)

//go:embed web/assets/* web/ui-templates/* web/static-templates/*
var assets embed.FS

func main() {
	cmd.Embedded = assets
	cmd.Execute()
}
