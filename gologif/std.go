// Copyright 2020 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gologif

import (
	"io"
	"os"

	"github.com/shimt/go-logif"
)

var std = New(os.Stderr, "", LstdFlags)

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return std.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	std.p(v)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.pf(format, v)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.pl(v)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.p(v)
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	std.pf(format, v)
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	std.pl(v)
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	panic(std.p(v))
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	panic(std.pf(format, v))
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	panic(std.pl(v))
}

// Debug write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	if std.OutputLevel() > logif.DEBUG {
		return
	}

	std.lp(logif.DEBUG, v)
}

// Debugf write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	if std.OutputLevel() > logif.DEBUG {
		return
	}

	std.lpf(logif.DEBUG, format, v)
}

// Debugln write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...interface{}) {
	if std.OutputLevel() > logif.DEBUG {
		return
	}

	std.lpl(logif.DEBUG, v)
}

// Info write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	if std.OutputLevel() > logif.INFO {
		return
	}

	std.lp(logif.INFO, v)
}

// Infof write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	if std.OutputLevel() > logif.INFO {
		return
	}

	std.lpf(logif.INFO, format, v)
}

// Infoln write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Println.
func Infoln(v ...interface{}) {
	if std.OutputLevel() > logif.INFO {
		return
	}

	std.lpl(logif.INFO, v)
}

// Warn write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {
	if std.OutputLevel() > logif.WARN {
		return
	}

	std.lp(logif.WARN, v)
}

// Warnf write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	if std.OutputLevel() > logif.WARN {
		return
	}

	std.lpf(logif.WARN, format, v)
}

// Warnln write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Println.
func Warnln(v ...interface{}) {
	if std.OutputLevel() > logif.WARN {
		return
	}

	std.lpl(logif.WARN, v)
}

// Error write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	if std.OutputLevel() > logif.ERROR {
		return
	}

	std.lp(logif.ERROR, v)
}

// Errorf write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	if std.OutputLevel() > logif.ERROR {
		return
	}

	std.lpf(logif.ERROR, format, v)
}

// Errorln write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Println.
func Errorln(v ...interface{}) {
	if std.OutputLevel() > logif.ERROR {
		return
	}

	std.lpl(logif.ERROR, v)
}

// SetOutputLevel set output level
func SetOutputLevel(l logif.LogLevel) {
	std.SetOutputLevel(l)
}

// OutputLevel set output level
func OutputLevel() logif.LogLevel {
	return std.OutputLevel()
}
