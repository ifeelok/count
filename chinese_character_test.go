package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountChineseCharacter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty string",
			args: args{""},
			want: 0,
		},
		{
			name: "string with no chinese characters",
			args: args{"123"},
			want: 0,
		},
		{
			name: "string with chinese characters",
			args: args{"你好"},
			want: 2,
		},
		{
			name: "string with chinese characters and other characters",
			args: args{"你好123"},
			want: 2,
		},
		{
			name: "string with only other characters",
			args: args{"!@#$%^&*()"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountChineseCharacter(tt.args.s); got != tt.want {
				t.Errorf("CountChineseCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChineseCharacterCount(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "empty string",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "only english",
			input: "Hello, world!",
			want:  map[string]int{},
		},
		{
			name:  "only chinese",
			input: "你好，世界！",
			want: map[string]int{
				"你": 1,
				"好": 1,
				"，": 1,
				"世": 1,
				"界": 1,
				"！": 1,
			},
		},
		{
			name:  "english and chinese",
			input: "Hello, 你好，世界！",
			want: map[string]int{
				"你": 1,
				"好": 1,
				"，": 1,
				"世": 1,
				"界": 1,
				"！": 1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := GetChineseCharacterCount(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
