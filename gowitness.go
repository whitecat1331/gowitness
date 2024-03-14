package gowitness

import (
	"embed"

	"github.com/sensepost/gowitness/cmd"
)

//go:embed web/assets/* web/ui-templates/* web/static-templates/*
var assets embed.FS

func GoWitness() {
	cmd.Embedded = assets
	cmd.Execute()
}
