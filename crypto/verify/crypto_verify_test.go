package verify

import (
	"bytes"
	"github.com/google/gofuzz"
	"testing"
)

var testCount = 100000

func TestVerify(t *testing.T) {
	// Fuzzing
	f := fuzz.New()

	// Run tests
	for i := 0; i < testCount; i++ {
		var x, y [64]byte

		// Create random x & y
		f.Fuzz(&x)
		f.Fuzz(&y)

		// Check 16
		if Verify16(x[:16], y[:16]) != bytes.Equal(x[:16], y[:16]) {
			t.Errorf("Verify16: invalid verification result %v", Verify16(x[:16], y[:16]))
		}

		// Check 32
		if Verify32(x[:32], y[:32]) != bytes.Equal(x[:32], y[:32]) {
			t.Errorf("Verify32: invalid verification result %v", Verify32(x[:32], y[:32]))
		}

		// Check 64
		if Verify64(x[:64], y[:64]) != bytes.Equal(x[:64], y[:64]) {
			t.Errorf("Verify64: invalid verification result %v", Verify64(x[:64], y[:64]))
		}
	}
}
