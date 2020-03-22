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

func Test_Logger_Debug(t *testing.T) {
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

			l.SetOutputLevel(DEBUG)
			l.Debug(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("Logger.Debug = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Debugf(t *testing.T) {
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

			l.SetOutputLevel(DEBUG)
			l.Debugf(tt.args.f, tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("Logger.Debugf = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Debugln(t *testing.T) {
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

			l.SetOutputLevel(DEBUG)
			l.Debugln(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("Logger.Debugln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Info(t *testing.T) {
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

			l.SetOutputLevel(INFO)
			l.Info(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("Logger.Info = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Infof(t *testing.T) {
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
				t.Errorf("Logger.Infof = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Infoln(t *testing.T) {
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

			l.SetOutputLevel(INFO)
			l.Infoln(tt.args.v...)

			if got := b.String(); !strings.HasSuffix(got, tt.want) {
				t.Errorf("Logger.Infoln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Warn(t *testing.T) {
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
				t.Errorf("Logger.Warn = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Warnf(t *testing.T) {
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
				t.Errorf("Logger.Warnf = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Warnln(t *testing.T) {
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
				t.Errorf("Logger.Warnln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Error(t *testing.T) {
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
				t.Errorf("Logger.Error = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Errorf(t *testing.T) {
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
				t.Errorf("Logger.Errorf = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_Errorln(t *testing.T) {
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
				t.Errorf("Logger.Errorln = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Logger_calldepth(t *testing.T) {
	want := "gologif_test.go:"
	b := &bytes.Buffer{}
	l := New(b, "", Lshortfile)
	l.Print("test")

	got := b.String()
	if !strings.HasPrefix(got, want) {
		t.Errorf("got = %v, want prefix %v", got, want)
	}
}
