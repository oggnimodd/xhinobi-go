package cli

import (
	"github.com/spf13/cobra"
)

var Flags struct {
	PrependFileName bool
	Minify          bool
	IgnorePatterns  []string
	WithTree        bool
	OSC52           bool
}

var RootCmd *cobra.Command

func SetupRootCommand() {
	RootCmd = &cobra.Command{
		Use: "xhinobi",
	}
	RootCmd.Flags().BoolVarP(&Flags.PrependFileName, "prependFileName", "n", false, "Prepend the file name before the content")
	RootCmd.Flags().BoolVarP(&Flags.Minify, "minify", "m", false, "Minify the output")
	RootCmd.Flags().StringSliceVarP(&Flags.IgnorePatterns, "ignore", "i", []string{}, "Glob patterns to ignore (can be used multiple times)")
	RootCmd.Flags().BoolVarP(&Flags.WithTree, "tree", "t", false, "Prepend the output with a directory tree (requires 'tree' command)")
	RootCmd.Flags().BoolVarP(&Flags.OSC52, "osc52", "o", false, "Use OSC52 escape sequence for clipboard over SSH")
}
