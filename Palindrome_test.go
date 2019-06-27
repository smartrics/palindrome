package palindrome

import "testing"

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
