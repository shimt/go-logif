// Copyright 2016 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logif is logging interface
package logif

import "io"

//go:generate stringer -type=LogLevel

// LogLevel is log message level
type LogLevel int32

const (
	// DEBUG debug log level
	DEBUG LogLevel = iota
	// MINLEVEL min log level
	MINLEVEL LogLevel = iota - 1
	// INFO info log level
	INFO LogLevel = iota
	// WARN warnning log level
	WARN LogLevel = iota
	// ERROR error log level
	ERROR LogLevel = iota
	// FATAL fatal log level
	FATAL LogLevel = iota
	// PANIC panic log level
	PANIC LogLevel = iota
	// MAXLEVEL max log level
	MAXLEVEL LogLevel = iota - 1
)

// GoLogger minimum logging interface
type GoLogger interface {
	// Output writes the output for a logging event.
	// The string s contains the text to print after the prefix specified
	// by the flags of the Logger. A newline is appended if the last character
	// of s is not already a newline.
	// Calldepth is used to recover the PC and is provided for generality,
	// although at the moment on all pre-defined paths it will be 2
	Output(calldepth int, s string) error

	// Print calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Print.
	Print(v ...interface{})
	// Printf calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})
	// Println calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Println.
	Println(v ...interface{})

	// Fatal write message(level=ERROR) to the logger followed by a call to os.Exit(1).
	// Arguments are handled in the manner of fmt.Print.
	Fatal(v ...interface{})
	// Fatalf write message(level=ERROR) to the logger followed by a call to os.Exit(1).
	// Arguments are handled in the manner of fmt.Printf.
	Fatalf(format string, v ...interface{})
	// Fatalln iwrite message(level=ERROR) to the logger followed by a call to os.Exit(1).
	// Arguments are handled in the manner of fmt.Println.
	Fatalln(v ...interface{})

	// Panic write message(level=PANIC) to the logger followed by a call to panic().
	// Arguments are handled in the manner of fmt.Print.
	Panic(v ...interface{})
	// Panicf write message(level=PANIC) to the logger followed by a call to panic().
	// Arguments are handled in the manner of fmt.Printf.
	Panicf(format string, v ...interface{})
	// Panicln write message(level=PANIC) to the logger followed by a call to panic().
	// Arguments are handled in the manner of fmt.Println.
	Panicln(v ...interface{})
}

//GoLoggerModifier leveld logging modifier interface
type GoLoggerModifier interface {
	// SetFlags sets the output flags for the logger.
	SetFlags(flag int)

	// Flags returns the output flags for the logger.
	Flags()

	// SetPrefix sets the output prefix for the logger.
	SetPrefix(prefix string)

	// Prefix returns the output prefix for the logger.
	Prefix()

	// SetOutput sets the output destination for the logger.
	SetOutput(w io.Writer)
}

//LeveledLogger leveld logging interface
type LeveledLogger interface {
	// Debug write message(level=DEBUG) to the logger.
	// Arguments are handled in the manner of fmt.Print.
	Debug(v ...interface{})
	// Debugf write message(level=DEBUG) to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Debugf(format string, v ...interface{})
	// Debugln write message(level=DEBUG) to the logger.
	// Arguments are handled in the manner of fmt.Println.
	Debugln(v ...interface{})

	// Info write message(level=INFO) to the logger.
	// Arguments are handled in the manner of fmt.Print.
	Info(v ...interface{})
	// Infof write message(level=INFO) to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Infof(format string, v ...interface{})
	// Infoln write message(level=INFO) to the logger.
	// Arguments are handled in the manner of fmt.Println.
	Infoln(v ...interface{})

	// Warn write message(level=WARN) to the logger.
	// Arguments are handled in the manner of fmt.Print.
	Warn(v ...interface{})
	// Warnf write message(level=WARN) to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Warnf(format string, v ...interface{})
	// Warnln write message(level=WARN) to the logger.
	// Arguments are handled in the manner of fmt.Println.
	Warnln(v ...interface{})

	// Error write message(level=ERROR) to the logger.
	// Arguments are handled in the manner of fmt.Print.
	Error(v ...interface{})
	// Errorf write message(level=ERROR) to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, v ...interface{})
	// Errorln write message(level=ERROR) to the logger.
	// Arguments are handled in the manner of fmt.Println.
	Errorln(v ...interface{})
}

//LeveledLoggerModifier leveld logging modifier interface
type LeveledLoggerModifier interface {
	// SetOutputLevel set output level
	SetOutputLevel(l LogLevel)

	// OutputLevel set output level
	OutputLevel() LogLevel
}

// Logger logging interface
type Logger interface {
	GoLogger
	LeveledLogger
}

// ToGoLogger cast to LevelLogger interface
func ToGoLogger(v interface{}) (l GoLogger, ok bool) {
	l, ok = v.(GoLogger)
	return
}

// ToGoLoggerModifier cast to ToGoLoggerModifier interface
func ToGoLoggerModifier(v interface{}) (l GoLoggerModifier, ok bool) {
	l, ok = v.(GoLoggerModifier)
	return
}

// ToLeveledLogger cast to LevelLogger interface
func ToLeveledLogger(v interface{}) (l LeveledLogger, ok bool) {
	l, ok = v.(LeveledLogger)
	return
}

// ToLeveledLoggerModifier cast to ToGoLoggerModifier interface
func ToLeveledLoggerModifier(v interface{}) (l LeveledLoggerModifier, ok bool) {
	l, ok = v.(LeveledLoggerModifier)
	return
}

// ToLogger cast to ToLogger interface
func ToLogger(v interface{}) (l Logger, ok bool) {
	l, ok = v.(Logger)
	return
}
