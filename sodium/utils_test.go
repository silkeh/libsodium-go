package sodium

import (
	"bytes"
	"fmt"
	"github.com/google/gofuzz"
	"math/big"
	"testing"
)

// ReverseInPlace reverses the data in `b` in place.
func ReverseInplace(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// Reverse returns the data in `a` in reverse order.
func Reverse(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	ReverseInplace(b)
	return b
}

func FixSize(a []byte, length int) []byte {
	dif := length - len(a)
	if len(a) == 0 {
		a = nil
	} else if dif > 0 {
		a = append(make([]byte, dif), a...)
	} else if dif < 0 {
		a = a[-dif:]
	}

	return a
}

func TestSodiumUtils(t *testing.T) {
	// Load fuzzing
	f := fuzz.New().NumElements(1, 1024)

	for i := 0; i < 10000; i++ {
		var a, b []byte
		f.Fuzz(&a)
		f.Fuzz(&b)

		// Get the minimum length of the two slices
		length := len(a)
		if len(a) > len(b) {
			length = len(b)
		}

		// Make the slices the same length
		a = a[:length]
		b = b[:length]

		// Test MemZero
		z := make([]byte, len(a))
		copy(z, a)
		MemZero(z)
		if !bytes.Equal(z, make([]byte, length)) {
			t.Error("MemZero failed to set bytes to zero.")
			t.FailNow()
		}

		// Test MemCmp
		sb := MemCmp(a, b)
		gb := bytes.Equal(a, b)
		if sb != gb {
			t.Errorf("MemCmp failed (%v != %v) for: %x, %x", sb, gb, a, b)
			t.FailNow()
		}

		// Test Reverse
		ra := Reverse(a)
		rb := Reverse(b)
		if !bytes.Equal(Reverse(ra), a) {
			t.Error("Reverse of reverse is not the same as original.")
			t.FailNow()
		}

		// Test Compare
		si := Compare(ra, rb)
		gi := bytes.Compare(a, b)
		if si != gi {
			t.Errorf("Compare failed %v should be %v for: %x, %x", si, gi, a, b)
			t.FailNow()
		}

		// Test IsZero
		if IsZero(a) != bytes.Equal(a, make([]byte, len(a))) {
			t.Errorf("IsZero failed for %x", a)
			t.FailNow()
		}

		// Increment bytes via big.Int
		bi := new(big.Int).SetBytes(a)
		ap := bi.Add(bi, big.NewInt(1)).Bytes()
		ap = FixSize(ap, len(a))

		// Test increment
		ra = Reverse(a)
		Increment(ra)
		if len(a) > 0 && !bytes.Equal(Reverse(ra), ap) {
			t.Errorf("Increment failed for %x", a)
			t.FailNow()
		}

		// Add a and b via big.Int
		ba, bb := new(big.Int).SetBytes(a), new(big.Int).SetBytes(b)
		ap = new(big.Int).Add(ba, bb).Bytes()
		ap = FixSize(ap, len(a))

		// Test add
		ra = Reverse(a)
		Add(ra, rb)
		if !bytes.Equal(Reverse(ra), ap) {
			t.Errorf("Addition failed for %#x + %#x", a, b)
			t.FailNow()
		}

		// Test Bin2Hex
		sHex := Bin2Hex(a)
		gHex := fmt.Sprintf("%x", a)
		if sHex != gHex {
			t.Errorf("Bin2Hex failed: %s != %s", sHex, gHex)
			t.FailNow()
		}

		// Encode with spaces and capitals
		gHex = fmt.Sprintf("% X", a)

		// Test Hex2Bin
		ha := Hex2Bin(gHex, " ")
		if !bytes.Equal(ha, a) {
			t.Errorf("Hex2Bin failed for %s: %x != %x", gHex, ha, a)
			t.FailNow()
		}
	}
}
