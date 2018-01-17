// Copyright 2016 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logrusif is logif wrapper for logrus package
package logrusif

import (
	"sync/atomic"

	"github.com/Sirupsen/logrus"
	logif "github.com/shimt/go-logif"
)

type logger struct {
	entity      *logrus.Logger
	outputLevel int32
}

var level2logrus []logrus.Level

func init() {
	level2logrus = make([]logrus.Level, logif.MAXLEVEL+1)
	level2logrus[logif.DEBUG] = logrus.DebugLevel
	level2logrus[logif.INFO] = logrus.InfoLevel
	level2logrus[logif.WARN] = logrus.WarnLevel
	level2logrus[logif.ERROR] = logrus.ErrorLevel
}

// Print calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Print.
func (l *logger) Print(v ...interface{}) {
	l.entity.Print(v...)
}

// Printf calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Printf(format string, v ...interface{}) {
	l.entity.Printf(format, v...)
}

// Println calls l.Output to print to the logger. Arguments are handled in the manner of fmt.Println.
func (l *logger) Println(v ...interface{}) {
	l.entity.Println(v...)
}

// Fatal write message(level=ERROR) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Fatal(v ...interface{}) {
	l.entity.Fatal(v...)
}

// Fatalf write message(level=ERROR) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Fatalf(format string, v ...interface{}) {
	l.entity.Fatalf(format, v...)
}

// Fatalln iwrite message(level=ERROR) to the logger followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Fatalln(v ...interface{}) {
	l.entity.Fatalln(v...)
}

// Panic write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Panic(v ...interface{}) {
	l.entity.Panic(v...)
}

// Panicf write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Panicf(format string, v ...interface{}) {
	l.entity.Panicf(format, v...)
}

// Panic write message(level=PANIC) to the logger followed by a call to panic().
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Panicln(v ...interface{}) {
	l.entity.Panicln(v...)
}

// Debug write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Debug(v ...interface{}) {
	l.entity.Debug(v...)
}

// Debugf write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(format string, v ...interface{}) {
	l.entity.Debugf(format, v...)
}

// Debugln write message(level=DEBUG) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Debugln(v ...interface{}) {
	l.entity.Debugln(v...)
}

// Info write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Info(v ...interface{}) {
	l.entity.Info(v...)
}

// Infof write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Infof(format string, v ...interface{}) {
	l.entity.Infof(format, v...)
}

// Infoln write message(level=INFO) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Infoln(v ...interface{}) {
	l.entity.Infoln(v...)
}

// Warn write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Warn(v ...interface{}) {
	l.entity.Warn(v...)
}

// Warnf write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Warnf(format string, v ...interface{}) {
	l.entity.Warnf(format, v...)
}

// Warnln write message(level=WARN) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Warnln(v ...interface{}) {
	l.entity.Warnln(v...)
}

// Error write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Print.
func (l *logger) Error(v ...interface{}) {
	l.entity.Error(v...)
}

// Errorf write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(format string, v ...interface{}) {
	l.entity.Errorf(format, v...)
}

// Errorln write message(level=ERROR) to the logger.
// Arguments are handled in the manner of fmt.Println.
func (l *logger) Errorln(v ...interface{}) {
	l.entity.Errorln(v...)
}

// SetOutputLevel set output level
func (l *logger) SetOutputLevel(level logif.LogLevel) {
	atomic.StoreInt32(&l.outputLevel, int32(level))
	l.entity.SetLevel(level2logrus[level])
}

// OutputLevel set output level
func (l *logger) OutputLevel() logif.LogLevel {
	return logif.LogLevel(atomic.LoadInt32(&l.outputLevel))
}

// ToGoLogger cast to logif.GoLogger
func (l *logger) ToGoLogger() logif.GoLogger {
	return l.entity
}

// ToLeveledLogger cast to logif.LeveledLogger
func (l *logger) ToLeveledLogger() logif.LeveledLogger {
	return l.entity
}

// New create new logger instance.
func New() (l logif.Logger) {
	return &logger{
		entity:      logrus.New(),
		outputLevel: int32(logif.WARN),
	}
}
