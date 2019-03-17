package owiki

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "owiki",
		Short: "A simple web file editor to test golang stuff.",
		Long:  "A simple web file editor to test golang stuff; adapted from the gowiki example.",
	}
)
