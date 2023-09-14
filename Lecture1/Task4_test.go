package Lecture1

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test not sorted array",
			args: args{[]int{4, 6, 12, 77, 2, 3, 5, 88}},
			want: []int{2, 3, 4, 5, 6, 12, 77, 88},
		},
		{
			name: "Test  sorted array",
			args: args{[]int{2, 3, 4, 5, 6, 12, 77, 88}},
			want: []int{2, 3, 4, 5, 6, 12, 77, 88},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortWithGoLibrary(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test not sorted array",
			args: args{[]int{4, 6, 12, 77, 2, 3, 5, 88}},
			want: []int{2, 3, 4, 5, 6, 12, 77, 88},
		},
		{
			name: "Test  sorted array",
			args: args{[]int{2, 3, 4, 5, 6, 12, 77, 88}},
			want: []int{2, 3, 4, 5, 6, 12, 77, 88},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortWithGoLibrary(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortWithGoLibrary() = %v, want %v", got, tt.want)
			}
		})
	}
}
