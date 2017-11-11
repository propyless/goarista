// Copyright (C) 2017  Arista Networks, Inc.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

const (
	goldFile = "testdata/test.go.gold"
	inFile   = "testdata/test.go.in"
)

func TestImportSort(t *testing.T) {
	in, err := ioutil.ReadFile(inFile)
	if err != nil {
		t.Fatal(err)
	}
	gold, err := ioutil.ReadFile(goldFile)
	if err != nil {
		t.Fatal(err)
	}
	sections := []string{"foobar", "cvshub.com/foobar"}
	if out, err := genFile(gold, sections); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal(out, gold) {
		t.Error("importsort on test.go.gold file produced a change")
		t.Log(string(out))
	}
	if out, err := genFile(in, sections); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal(out, gold) {
		t.Error("importsort on test.go.in different than gold")
		t.Log(string(out))
	}
}
