package count

import (
	"reflect"
	"testing"
)

func TestGetEnglishWordCount(t *testing.T) {
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
			name:  "one word",
			input: "hello",
			want:  map[string]int{"hello": 1},
		},
		{
			name:  "multiple same word",
			input: "hello hello",
			want:  map[string]int{"hello": 2},
		},
		{
			name:  "multiple different words",
			input: "hello world",
			want:  map[string]int{"hello": 1, "world": 1},
		},
		{
			name:  "multiple words with punctuation",
			input: "hello, world!",
			want:  map[string]int{"hello": 1, "world": 1},
		},
		{
			name:  "multiple words with numbers",
			input: "hello, world123",
			want:  map[string]int{"hello": 1, "world": 1},
		},
		{
			name:  "complex sentence",
			input: "This is a complex sentence for testing.",
			want:  map[string]int{"this": 1, "is": 1, "a": 1, "complex": 1, "sentence": 1, "for": 1, "testing": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnglishWordCount(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnglishWordCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountEnglishWord(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test1",
			input: "Hello world! This is a test.",
			want:  6,
		},
		{
			name:  "test2",
			input: "Hello, world! This is a test.",
			want:  6,
		},
		{
			name:  "test3",
			input: "Hello world! This is a test.",
			want:  6,
		},
		{
			name:  "test4",
			input: "Hello, world! This is a test.",
			want:  6,
		},
		{
			name:  "test5",
			input: "Hello, world! This is a test.",
			want:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountEnglishWord(tt.input); got != tt.want {
				t.Errorf("CountEnglishWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
