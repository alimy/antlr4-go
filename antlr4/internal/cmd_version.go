// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package internal

import (
	"fmt"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var (
	appVer = semver.Version{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}

	antlrVer = semver.Version{
		Major: 4,
		Minor: 9,
		Patch: 0,
	}
)

func init() {
	register(&cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	})
}

func versionRun(_cmd *cobra.Command, _args []string) {
	fmt.Printf("%s\nantlr4 %s\n", appVer, antlrVer)
}
