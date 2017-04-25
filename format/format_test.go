package format

import (
	"testing"
	"time"
)

func TestToInt(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.v); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.v); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeFromTi(t *testing.T) {
	type args struct {
		template  string
		timepoint time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeFromTi(tt.args.template, tt.args.timepoint); got != tt.want {
				t.Errorf("FormatTimeFromTi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeFromUnix(t *testing.T) {
	type args struct {
		template  string
		timepoint int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeFromUnix(tt.args.template, tt.args.timepoint); got != tt.want {
				t.Errorf("FormatTimeFromUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToUnix(t *testing.T) {
	type args struct {
		timepoint string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTimeToUnix(tt.args.timepoint); got != tt.want {
				t.Errorf("FormatTimeToUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}
