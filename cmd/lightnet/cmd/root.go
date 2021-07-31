package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "lightnet",
	Short: "an CLI tool for lightnet",
	Long:  "an CLI tool providing interactive ways to operate lightnet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run lightnet")
	},
}

func Execute(in io.Reader, out, errOut io.Writer) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
