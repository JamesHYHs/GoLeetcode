package problems

import (
	"testing"
)

func Test_searchInsertQ35(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsertQ35(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsertQ35() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxAverageQ643(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverageQ643(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverageQ643() = %v, want %v", got, tt.want)
			}
		})
	}
}
