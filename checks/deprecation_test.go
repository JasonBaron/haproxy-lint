package checks

import (
	"testing"

	"github.com/abulimov/haproxy-lint/lib"
)

func TestCheckDeprecations(t *testing.T) {
	lines, err := lib.ReadConfigFile("../testdata/haproxy.cfg")
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}
	sections := lib.GetSection("frontend", lines)
	problems := CheckDeprecations(sections[1])

	if len(problems) != 3 {
		t.Errorf("Expected %d problems, got %d", 3, len(problems))
	}
}
