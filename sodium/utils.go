package sodium

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"
import (
	"github.com/GoKillers/libsodium-go/support"
	"unsafe"
)

// MemZero erases the memory in buffer `buff1`.
func MemZero(buff1 []byte) {
	if len(buff1) > 0 {
		C.sodium_memzero(
			unsafe.Pointer(&buff1[0]),
			C.size_t(len(buff1)))
	}
}

// MemCmp checks if the data in `buff1` and `buff2` is identical.
// This function is performs in a constant time to avoid leaking information.
func MemCmp(buff1, buff2 []byte) bool {
	support.CheckSizeEqual(buff1, buff2, "buffer 1", "buffer 2")

	return int(C.sodium_memcmp(
		unsafe.Pointer(support.BytePointer(buff1)),
		unsafe.Pointer(support.BytePointer(buff2)),
		C.size_t(len(buff1)))) == 0
}

// Compare compares the data in `a` and `b` in a constant time.
// It returns 0 if a==b, -1 if a < b, and 1 if a > b.
// Note: this comparison is little-endian.
func Compare(a, b []byte) int {
	support.CheckSizeEqual(a, b, "a", "b")
	C.sodium_init()
	return int(C.sodium_compare(
		(*C.uchar)(support.BytePointer(a)),
		(*C.uchar)(support.BytePointer(b)),
		C.size_t(len(a))))
}

// IsZero checks if the data in `n` is all zeroes in a constant time.
func IsZero(n []byte) bool {
	exit := int(C.sodium_is_zero(
		(*C.uchar)(support.BytePointer(n)),
		C.size_t(len(n))))

	return exit == 1
}

// Increment increments a large number `n` in place in constant time.
// Note: this increment is little-endian.
func Increment(n []byte) {
	C.sodium_increment(
		(*C.uchar)(support.BytePointer(n)),
		C.size_t(len(n)))
}

// Add adds two large numbers `a` and `b` in constant time,
// and overwrites `a` with the result.
func Add(a, b []byte) {
	support.CheckSizeEqual(a, b, "a", "b")
	C.sodium_add(
		(*C.uchar)(support.BytePointer(a)),
		(*C.uchar)(support.BytePointer(b)),
		C.size_t(len(a)))
}

// Bin2Hex returns the hexadecimal string representation of binary data `bin`.
func Bin2Hex(bin []byte) string {
	hex := make([]int8, len(bin)*2+1)

	ret := C.sodium_bin2hex(
		(*C.char)(&hex[0]),
		C.size_t(len(hex)),
		(*C.uchar)(support.BytePointer(bin)),
		C.size_t(len(bin)))

	return C.GoString(ret)
}

// Hex2Bin decodes a hexadecimal string `hex` into binary data.
// Characters in `ignore` will be ignored.
func Hex2Bin(hex, ignore string) []byte {
	bin := make([]byte, len(hex)/2)
	var length C.size_t

	C.sodium_hex2bin(
		(*C.uchar)(support.BytePointer(bin)),
		C.size_t(len(bin)),
		(*C.char)(C.CString(hex)),
		C.size_t(len(hex)),
		(*C.char)(C.CString(ignore)),
		(*C.size_t)(&length),
		(**C.char)(nil))

	return bin[:int(length)]
}
