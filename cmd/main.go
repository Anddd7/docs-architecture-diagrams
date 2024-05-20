package main

import (
	"log/slog"

	"github.com/alecthomas/kong"
)

type Globals struct {
}

type CLI struct {
	Globals

	Translate TranslateCmd `cmd:"" help:"Translate the excalidraw files."`
}

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("arch"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	err := ctx.Run(&cli.Globals)
	if err != nil {
		slog.Error("failed", "err", err)
	}
}
