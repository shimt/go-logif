// Copyright 2016 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gologif is wrapper of a standard "log" package.
package gologif

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"

	"github.com/shimt/go-logif"
)

var levelString []string
var levelStringWithSpace []string

func init() {
	var l logif.LogLevel
	levelString = make([]string, logif.MAXLEVEL+1)
	levelStringWithSpace = make([]string, logif.MAXLEVEL+1)
	for l = logif.MINLEVEL; l <= logif.MAXLEVEL; l++ {
		levelString[l] = "[" + l.String() + "]"
		levelStringWithSpace[l] = "[" + l.String() + "] "
	}
}

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	Lmsgprefix    = log.Lmsgprefix
	LstdFlags     = log.LstdFlags
)

const (
	DEBUG = logif.DEBUG
	INFO  = logif.INFO
	WARN  = logif.WARN
	ERROR = logif.ERROR
)

// Logger is wrapper for Golang default logger (log.Logger).
type Logger struct {
	entity      *log.Logger
	outputLevel int32
}

// verify interface compliance.
var (
	_ logif.Logger                = (*Logger)(nil)
	_ logif.LoggerModifier        = (*Logger)(nil)
	_ logif.LeveledLogger         = (*Logger)(nil)
	_ logif.LeveledLoggerModifier = (*Logger)(nil)

	_ logif.Logger         = (*log.Logger)(nil)
	_ logif.LoggerModifier = (*log.Logger)(nil)
)

// SetFlags sets the output flags for the logger.
func (l *Logger) SetFlags(flag int) {
	l.entity.SetFlags(flag)
}

// Flags returns the output flags for the logger.
func (l *Logger) Flags() int {
	return l.entity.Flags()
}

// SetPrefix sets the output prefix for the logger.
func (l *Logger) SetPrefix(prefix string) {
	l.entity.SetPrefix(prefix)
}

// Prefix returns the output prefix for the logger.
func (l *Logger) Prefix() string {
	return l.entity.Prefix()
}

// Writer returns the output destination for the logger.
func (l *Logger) Writer() io.Writer {
	return l.entity.Writer()
}

// Output writes the output for a logging event.
// The string s contains the text to print after the prefix specified
// by the flags of the Logger. A newline is appended if the last character
// of s is not already a newline.
// Calldepth is used to recover the PC and is provided for generality,
// although at the moment on all pre-defined paths it will be 4
func (l *Logger) Output(calldepth int, s string) error {
	return l.entity.Output(calldepth, s)
}

// SetOutput sets the output destination for the logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.entity.SetOutput(w)
}

func (l *Logger) p(v []interface{}) string {
	s := fmt.Sprint(v...)
	l.Output(4, s)
	return s
}

func (l *Logger) pf(f string, v []interface{}) string {
	s := fmt.Sprintf(f, v...)
	l.Output(4, s)
	return s
}

func (l *Logger) pl(v []interface{}) string {
	s := fmt.Sprintln(v...)
	l.Output(4, s)
	return s
}

// Print calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	l.p(v)
}

// Printf calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.pf(format, v)
}

// Println calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) {
	l.pl(v)
}

// Fatal write message(level=FATAL) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Fatal(v ...interface{}) {
	l.p(v)
	os.Exit(1)
}

// Fatalf write message(level=FATAL) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.pf(format, v)
	os.Exit(1)
}

// Fatalln iwrite message(level=FATAL) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Fatalln(v ...interface{}) {
	l.pl(v)
	os.Exit(1)
}

// Panic write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Panic(v ...interface{}) {
	panic(l.p(v))
}

// Panicf write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Panicf(format string, v ...interface{}) {
	panic(l.pf(format, v))
}

// Panicln write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Panicln(v ...interface{}) {
	panic(l.pl(v))
}

func (l *Logger) lp(level logif.LogLevel, v []interface{}) {
	v2 := make([]interface{}, 0, len(v)+1)
	v2 = append(v2, levelStringWithSpace[level])
	v2 = append(v2, v...)
	l.Output(4, fmt.Sprint(v2...))
}

func (l *Logger) lpf(level logif.LogLevel, format string, v []interface{}) {
	l.Output(4, fmt.Sprintf(levelStringWithSpace[level]+format, v...))
}

func (l *Logger) lpl(level logif.LogLevel, v []interface{}) {
	v2 := make([]interface{}, 0, len(v)+1)
	v2 = append(v2, levelString[level])
	v2 = append(v2, v...)
	l.Output(4, fmt.Sprintln(v2...))
}

// Debug write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(v ...interface{}) {
	if l.OutputLevel() > logif.DEBUG {
		return
	}

	l.lp(logif.DEBUG, v)
}

// Debugf write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.OutputLevel() > logif.DEBUG {
		return
	}

	l.lpf(logif.DEBUG, format, v)
}

// Debugln write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Debugln(v ...interface{}) {
	if l.OutputLevel() > logif.DEBUG {
		return
	}

	l.lpl(logif.DEBUG, v)
}

// Info write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(v ...interface{}) {
	if l.OutputLevel() > logif.INFO {
		return
	}

	l.lp(logif.INFO, v)
}

// Infof write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.OutputLevel() > logif.INFO {
		return
	}

	l.lpf(logif.INFO, format, v)
}

// Infoln write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Infoln(v ...interface{}) {
	if l.OutputLevel() > logif.INFO {
		return
	}

	l.lpl(logif.INFO, v)
}

// Warn write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warn(v ...interface{}) {
	if l.OutputLevel() > logif.WARN {
		return
	}

	l.lp(logif.WARN, v)
}

// Warnf write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.OutputLevel() > logif.WARN {
		return
	}

	l.lpf(logif.WARN, format, v)
}

// Warnln write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warnln(v ...interface{}) {
	if l.OutputLevel() > logif.WARN {
		return
	}

	l.lpl(logif.WARN, v)
}

// Error write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(v ...interface{}) {
	if l.OutputLevel() > logif.ERROR {
		return
	}

	l.lp(logif.ERROR, v)
}

// Errorf write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.OutputLevel() > logif.ERROR {
		return
	}

	l.lpf(logif.ERROR, format, v)
}

// Errorln write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Errorln(v ...interface{}) {
	if l.OutputLevel() > logif.ERROR {
		return
	}

	l.lpl(logif.ERROR, v)
}

// SetOutputLevel set output level
func (l *Logger) SetOutputLevel(level logif.LogLevel) {
	atomic.StoreInt32(&l.outputLevel, int32(level))
}

// OutputLevel set output level
func (l *Logger) OutputLevel() logif.LogLevel {
	return logif.LogLevel(atomic.LoadInt32(&l.outputLevel))
}

// New create new logger instance.
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		entity:      log.New(out, prefix, flag),
		outputLevel: int32(logif.WARN),
	}
}
