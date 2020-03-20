// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gologif is wrapper of a standard "log" package.
package gologif

import (
	"bytes"
	"strings"
	"testing"
)

func Test_std_calldepth(t *testing.T) {
	want := "std_test.go:"
	b := &bytes.Buffer{}
	std.SetOutput(b)
	std.SetFlags(Lshortfile)
	std.Print("test")

	got := b.String()
	if !strings.HasPrefix(got, want) {
		t.Errorf("got = %v, want prefix %v", got, want)
	}
}
