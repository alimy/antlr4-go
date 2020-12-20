// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package main

import (
	"log"

	"github.com/alimy/antlr4-go/antlr4/internal"
)

func main() {
	if err := internal.Execute(); err != nil {
		log.Fatal(err)
	}
}
