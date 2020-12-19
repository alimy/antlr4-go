// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package main

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

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
