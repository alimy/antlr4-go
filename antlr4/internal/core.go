// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package internal

import (
	_ "embed"
	"os"
	"path/filepath"
)

const (
	antlrJarFile    = "antlr-4.9-complete.jar"
	tmpAntlrJarFile = "/tmp/antlr/antlr-4.9-complete.jar"
)

var (
	//go:embed jar/antlr-4.9-complete.jar
	jarRaw []byte

	antlrExec = &execRun{
		Describe: "antlr4 wraper",
		Cmd:      "/usr/bin/java",
		Argv: []string{
			"-jar",
			tmpAntlrJarFile,
		},
		Attr: defaultProcAttr(true),
	}
)

func writeAntlrJar(dstPath string, fileName string, force bool) error {
	path, err := filepath.EvalSymlinks(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(dstPath) {
				cwd, err := os.Getwd()
				if err != nil {
					return err
				}
				path = filepath.Join(cwd, dstPath)
			} else {
				path = dstPath
			}
		} else {
			return err
		}
	}

	filePath := filepath.Join(path, fileName)
	if existAntlrJar(filePath) && !force {
		return nil
	}

	dirPath := filepath.Dir(filePath)
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	if _, err = file.Write(jarRaw); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	return nil
}

func existAntlrJar(path string) bool {
	fi, err := os.Stat(path)
	if err != nil || fi.IsDir() {
		return false
	}
	return true
}

func preExec() error {
	return writeAntlrJar(filepath.Dir(tmpAntlrJarFile), filepath.Base(tmpAntlrJarFile), false)
}
