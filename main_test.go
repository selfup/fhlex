package main

import (
	"reflect"
	"testing"
)

func TestFormatBuf(t *testing.T) {
	type args struct {
		readBuf []byte
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FormatBuf(tt.args.readBuf)
			if got != tt.want {
				t.Errorf("FormatBuf() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FormatBuf() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreateBaudRateAndCmdData(t *testing.T) {
	tests := []struct {
		name  string
		want  uint
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateBaudRateAndCmdData()
			if got != tt.want {
				t.Errorf("CreateBaudRateAndCmdData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateBaudRateAndCmdData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
