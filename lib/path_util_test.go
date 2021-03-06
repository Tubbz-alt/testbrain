package lib

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCommonPathPrefix(t *testing.T) {
	t.Parallel()

	type testSampleModel struct {
		name     string
		input    []string
		expected string
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting working directory: %s", err)
	}

	testSamples := []testSampleModel{
		testSampleModel{
			name:     "Simple mismatch",
			input:    []string{"/a/b/c", "/a/b/d"},
			expected: "/a/b",
		},
		testSampleModel{
			name:     "All parts matching",
			input:    []string{"/a/b/c", "/a/b/c"},
			expected: "/a/b/c",
		},
		testSampleModel{
			name:     "Only one path",
			input:    []string{"/a/b/c"},
			expected: "/a/b/c",
		},
		testSampleModel{
			name:     "Matching a directory",
			input:    []string{"/a/b", "/a/b/c/d"},
			expected: "/a/b",
		},
		testSampleModel{
			name:     "Trailing directory",
			input:    []string{"/a/b", "/a/b/"},
			expected: "/a/b",
		},
		testSampleModel{
			name:     "No input should not crash",
			input:    []string{},
			expected: "",
		},
		testSampleModel{
			name:     "Multiple relative paths",
			input:    []string{"a/b/c", "a/b/d"},
			expected: filepath.Join(wd, "a/b"),
		},
	}

	for _, testSample := range testSamples {
		result, err := CommonPathPrefix(testSample.input)
		if err != nil {
			t.Fatalf("Unexpected error in test %s: %s", testSample.name, err)
		}
		if !reflect.DeepEqual(result, testSample.expected) {
			t.Fatalf("Unexpected results:\n\texpected: %v\n\tactual  : %v\n",
				testSample.expected, result)
		}
	}
}
