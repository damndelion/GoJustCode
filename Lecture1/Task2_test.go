package Lecture1

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Return right prefix",
			args: args{[]string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "Return empty string",
			args: args{[]string{"dog", "racecar", "car"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.s); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
