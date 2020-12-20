// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package internal

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	inDstPath string
	inForce   bool
)

func init() {
	jarCmd := &cobra.Command{
		Use:   "jar",
		Short: "inflate antlr jar file",
		Long:  "inflate antlr jar file",
		Run:   jarRun,
	}

	jarCmd.Flags().StringVarP(&inDstPath, "dst", "d", ".", "genereted destination target directory")
	jarCmd.Flags().BoolVarP(&inForce, "force", "f", false, "whether force generate")
	register(jarCmd)
}

func jarRun(_cmd *cobra.Command, _args []string) {
	if err := writeAntlrJar(inDstPath, antlrJarFile, inForce); err != nil {
		log.Fatal(err)
	}
}
