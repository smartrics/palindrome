package palindrome

import (
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty string is a palindrome", args{""}, true},
		{"Single char is a palindrome", args{"z"}, true},
		{"Non palindrome with odd len", args{"abcde"}, false},
		{"Palindrome with odd len", args{"abcba"}, true},
		{"Non palindrome with even len", args{"abcdef"}, false},
		{"Palindrome with eve len", args{"abccba"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.s); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeFrom(t *testing.T) {
	type args struct {
		s    string
		even bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty string makes an empty palindrome, even=true", args{"", true}, ""},
		{"Empty string makes an empty palindrome, even=false", args{"", false}, ""},
		{"Single char string makes 2 char string, even=true", args{"a", true}, "aa"},
		{"Single char string makes 1 char string, even=false", args{"a", false}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeFrom(tt.args.s, tt.args.even); got != tt.want {
				t.Errorf("MakeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
