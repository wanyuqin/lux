package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/wanyuqin/lux/app"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
		os.Exit(1)
	}
}
