package day02

import "testing"

func Test_findValidPasswordsFrom(t *testing.T) {
	type args struct {
		passwordPolicies []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Two Passwords are valid",
			args: args{
				passwordPolicies: []string{
					"1-3 a: abcde",
					"1-3 b: cdefg",
					"2-9 c: ccccccccc",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findValidPasswordsFrom(tt.args.passwordPolicies); got != tt.want {
				t.Errorf("findValidPasswordsFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
