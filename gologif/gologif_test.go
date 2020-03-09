// Copyright 2016 Shinichi MOTOKI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gologif

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/shimt/go-logif"
)

func Test_logger_Debug(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[DEBUG] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.SetOutputLevel(logif.DEBUG)
			l.Debug(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Debug = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Debugf(t *testing.T) {
	type args struct {
		f string
		v []interface{}
	}
	tests := []struct {
		name string

		args args
		want string
	}{
		{"string", args{f: "%s", v: []interface{}{"string"}}, "[DEBUG] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.SetOutputLevel(logif.DEBUG)
			l.Debugf(tt.args.f, tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Debugf = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Debugln(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[DEBUG] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.SetOutputLevel(logif.DEBUG)
			l.Debugln(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Debugln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Info(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[INFO] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.SetOutputLevel(logif.INFO)
			l.Info(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Info = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Infof(t *testing.T) {
	type args struct {
		f string
		v []interface{}
	}
	tests := []struct {
		name string

		args args
		want string
	}{
		{"string", args{f: "%s", v: []interface{}{"string"}}, "[INFO] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.SetOutputLevel(logif.INFO)
			l.Infof(tt.args.f, tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Infof = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Infoln(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[INFO] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.SetOutputLevel(logif.INFO)
			l.Infoln(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Infoln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Warn(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[WARN] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.Warn(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Warn = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Warnf(t *testing.T) {
	type args struct {
		f string
		v []interface{}
	}
	tests := []struct {
		name string

		args args
		want string
	}{
		{"string", args{f: "%s", v: []interface{}{"string"}}, "[WARN] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.Warnf(tt.args.f, tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Warnf = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Warnln(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[WARN] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.Warnln(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Warnln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Error(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[ERROR] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.Error(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Error = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Errorf(t *testing.T) {
	type args struct {
		f string
		v []interface{}
	}
	tests := []struct {
		name string

		args args
		want string
	}{
		{"string", args{f: "%s", v: []interface{}{"string"}}, "[ERROR] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.Errorf(tt.args.f, tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Errorf = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logger_Errorln(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{v: []interface{}{"string"}}, "[ERROR] string\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			l := New(b, "", log.LstdFlags)

			l.Errorln(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("logger.Errorln = %v, want %v", got, tt.want)
			}
		})
	}
}
