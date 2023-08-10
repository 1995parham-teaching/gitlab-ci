package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Nost", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("trad", pterm.NewStyle(pterm.FgLightRed)),
		putils.LettersFromStringWithStyle("amus", pterm.NewStyle(pterm.FgLightYellow)),
	).Render(); err != nil {
		_ = err
	}
}
