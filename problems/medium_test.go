package problems

import "testing"

func Test_lengthOfLongestSubstringQ3(t *testing.T) {
	type args struct {
		s string
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
				s: "abcabcbb",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringQ3(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringQ3() = %v, want %v", got, tt.want)
			}
		})
	}
}
