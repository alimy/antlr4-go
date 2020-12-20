// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package internal

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	inHelp bool
)

func init() {
	goCmd := &cobra.Command{
		Use:   "go",
		Short: "ANTLR Parser Generator for Golang",
		Long:  "ANTLR Parser Generator for Golang",
		Run:   goRun,
	}

	goCmd.Flags().BoolVarP(&inHelp, "help", "h", false, "show help info")
	register(goCmd)
}

func goRun(_cmd *cobra.Command, args []string) {
	if inHelp {
		log.Fatal(antlrExec.Run())
	}

	antlrGoExec := &execRun{
		Describe: "antlr4 wraper",
		Cmd:      "/usr/bin/java",
		Argv: []string{
			"-jar",
			tmpAntlrJarFile,
			"-Dlanguage=Go",
		},
		Attr: defaultProcAttr(true),
	}
	if err := preExec(); err != nil {
		log.Fatal(err)
	}
	if err := antlrGoExec.Run(args...); err != nil {
		log.Fatal(err)
	}
}
