package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd(in io.Reader, out, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lightnet",
		Short: "an CLI tool for lightnet",
		Long:  "an CLI tool providing interactive ways to operate lightnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Run lightnet")
			return nil
		},
	}
	cmd.SetIn(in)
	cmd.SetOut(out)
	cmd.SetErr(errOut)

	return cmd
}

func Execute(in io.Reader, out, errOut io.Writer) {
	rootCmd := newRootCmd(in, out, errOut)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
