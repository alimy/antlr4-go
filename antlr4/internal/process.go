// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this file is governed by the BSD 3-clause license that
// can be found in the LICENSE.txt file in the project root.

package internal

import (
	"errors"
	"os"
)

type execRun struct {
	Describe string
	Cmd      string
	Argv     []string
	Attr     *os.ProcAttr
}

func (r *execRun) Run(argv ...string) error {
	args := make([]string, 0, len(r.Argv)+len(argv))
	args = append(args, r.Argv...)
	args = append(args, argv...)

	r.checkProcAttr()
	process, err := os.StartProcess(r.Cmd, r.Argv, r.Attr)
	if err != nil {
		return err
	}
	ps, err := process.Wait()
	if err != nil {
		return err
	}
	if !ps.Success() {
		return errors.New(ps.String())
	}
	return nil
}

func (r *execRun) checkProcAttr() {
	if r.Attr == nil {
		r.Attr = &os.ProcAttr{}
		if homedir, err := os.UserHomeDir(); err == nil {
			r.Attr.Dir = homedir
		}
	}
}

func defaultProcAttr(inStdio bool) *os.ProcAttr {
	attr := &os.ProcAttr{}
	if inStdio {
		attr.Files = []*os.File{
			os.Stdin, os.Stdout, os.Stderr,
		}
	}
	if homedir, err := os.UserHomeDir(); err == nil {
		attr.Dir = homedir
	}
	return attr
}
