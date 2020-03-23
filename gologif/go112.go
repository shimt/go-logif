// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.12

package gologif

import "io"

// Writer returns the output destination for the logger.
func (l *Logger) Writer() io.Writer {
	return l.entity.Writer()
}

// Writer returns the output destination for the standard logger.
func Writer() io.Writer {
	return std.Writer()
}
