package styles

import "github.com/fatih/color"

// Style definitions for consistent output
var (
	Success = color.New(color.FgGreen, color.Bold)
	Error   = color.New(color.FgRed, color.Bold)
	Warning = color.New(color.FgYellow, color.Bold)
	Info    = color.New(color.FgCyan, color.Bold)
	Title   = color.New(color.FgMagenta, color.Bold, color.Underline)
)