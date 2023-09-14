package Lecture1

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Return correct indexes",
			args: args{[]int{1, 2, 3, 4}, 6},
			want: []int{1, 3},
		},
		{
			name: "Return empty slice",
			args: args{[]int{1, 2, 3, 4}, 9},
			want: []int{},
		},
		{
			name: "Return first matched indexes",
			args: args{[]int{1, 2, 3, 4, 5, 6}, 7},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
