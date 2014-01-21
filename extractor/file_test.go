package extractor

import (
	"testing"
)

func TestFileComplexity(t *testing.T) {
	complexity, err := FileComplexity("../example/example.go")

	if err != nil {
		t.Error("expected file to open, got ", err)
	}

	var expected int64 = 8
	if complexity != expected {
		t.Errorf("expected example file complexity to be %d, got %d", expected, complexity)
	}
}
