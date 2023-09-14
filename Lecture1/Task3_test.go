package Lecture1

import "testing"

func Test_compareTwoSlicesWithOrder(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test of true",
			args: args{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "Test of false",
			args: args{[]int{1, 2, 3, 3, 5}, []int{1, 2, 3, 4, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTwoSlicesWithOrder(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("compareTwoSlicesWithOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareTwoSlicesWithoutOrder(t *testing.T) {
	type args struct {
		s1 []int
		s2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test of true with order",
			args: args{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "Test of false with order",
			args: args{[]int{1, 2, 2, 4, 5}, []int{1, 2, 3, 4, 5}},
			want: false,
		},
		{
			name: "Test of true without order",
			args: args{[]int{1, 2, 3, 4, 5}, []int{1, 3, 4, 2, 5}},
			want: true,
		},
		{
			name: "Test of false with order",
			args: args{[]int{1, 2, 2, 4, 5}, []int{1, 2, 3, 1, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTwoSlicesWithoutOrder(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("compareTwoSlicesWithoutOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
