package jamfprointegration

import (
	"testing"
)

func Test_chooseMostAlphabeticalString(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "easy test - abc vs def and xyz",
			args: args{
				strings: []string{"abc", "def", "xyz"},
			},
			want: "abc",
		},
		{
			name: "ignore numbers",
			args: args{
				strings: []string{"123a", "jjj", "zzz"},
			},
			want: "123a",
		},
		{
			name: "hard test - azz vs rrs and byz",
			args: args{
				strings: []string{"byz", "rrs", "azz"},
			},
			want: "azz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chooseMostAlphabeticalString(tt.args.strings); got != tt.want {
				t.Errorf("chooseMostAlphabeticalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
