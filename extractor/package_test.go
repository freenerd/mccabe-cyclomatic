package extractor

import (
	"testing"
)

func TestPackageComplexity(t *testing.T) {
	complexity, err := PackageComplexity("github.com/freenerd/mccabe-cyclomatic/example")

	if err != nil {
		t.Error("expected package to open, got ", err)
	}

	var expected int64 = 8
	if complexity != expected {
		t.Errorf("expected example file complexity to be %d, got %d", expected, complexity)
	}
}
