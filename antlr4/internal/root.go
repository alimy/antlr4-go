package internal

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "antlr",
		Short: "ANTLR Parser Generator",
		Long:  "ANTLR Parser Generator",
		Run:   rootRun,
	}
)

func register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func rootRun(_cmd *cobra.Command, args []string) {
	var err error

	if err = preExec(); err == nil {
		err = antlrExec.Run(args...)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}
