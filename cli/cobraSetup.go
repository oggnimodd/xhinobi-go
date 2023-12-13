package cli

import (
	"github.com/spf13/cobra"
)

// Flags holds the flag values
var Flags struct {
	PrependFileName bool
	Minify          bool
}

// RootCmd is the root command for your application
var RootCmd *cobra.Command

func SetupRootCommand() {
	RootCmd = &cobra.Command{
		Use: "xhinobi",
	}

	RootCmd.Flags().BoolVarP(&Flags.PrependFileName, "prependFileName", "n", false, "Prepend the file name before the content")
	RootCmd.Flags().BoolVarP(&Flags.Minify, "minify", "m", false, "Minify the output")
}
