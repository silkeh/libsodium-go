package verify

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"
import "github.com/GoKillers/libsodium-go/support"

// The number of bytes for the various verifications
const (
	Bytes16 = C.crypto_verify_16_BYTES
	Bytes32 = C.crypto_verify_32_BYTES
	Bytes64 = C.crypto_verify_64_BYTES
)

// Verify16 returns true if inputs `x` and `y` are the same.
func Verify16(x, y []byte) bool {
	support.CheckSize(x, Bytes16, "x")
	support.CheckSize(y, Bytes16, "y")

	exit := C.crypto_verify_16(
		(*C.uchar)(&x[0]),
		(*C.uchar)(&y[0]))

	return exit == 0
}

// Verify32 returns true if inputs `x` and `y` are the same.
func Verify32(x, y []byte) bool {
	support.CheckSize(x, Bytes32, "x")
	support.CheckSize(y, Bytes32, "y")

	exit := C.crypto_verify_32(
		(*C.uchar)(&x[0]),
		(*C.uchar)(&y[0]))

	return exit == 0
}

// Verify64 returns true if inputs `x` and `y` are the same.
func Verify64(x, y []byte) bool {
	support.CheckSize(x, Bytes64, "x")
	support.CheckSize(y, Bytes64, "y")

	exit := C.crypto_verify_64(
		(*C.uchar)(&x[0]),
		(*C.uchar)(&y[0]))

	return exit == 0
}
