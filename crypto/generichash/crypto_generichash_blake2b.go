package generichash

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"
import (
	"github.com/silkeh/libsodium-go/support"
	"hash"
	"unsafe"
)

// The following constants reflect the properties of the Blake2b hash:
const (
	Blake2bBytesMin      int = C.crypto_generichash_blake2b_BYTES_MIN     // Minimum hash size in bytes
	Blake2bBytesMax      int = C.crypto_generichash_blake2b_BYTES_MAX     // Maximum hash size in bytes
	Blake2bBytes         int = C.crypto_generichash_blake2b_BYTES         // Default hash size in bytes
	Blake2bKeyBytesMin   int = C.crypto_generichash_blake2b_KEYBYTES_MIN  // Minimum key size in bytes
	Blake2bKeyBytesMax   int = C.crypto_generichash_blake2b_KEYBYTES_MAX  // Maximum key size in bytes
	Blake2bKeyBytes      int = C.crypto_generichash_blake2b_KEYBYTES      // Default key size in bytes
	Blake2bSaltBytes     int = C.crypto_generichash_blake2b_SALTBYTES     // Size of the salt in bytes
	Blake2bPersonalBytes int = C.crypto_generichash_blake2b_PERSONALBYTES // Size of the personal in bytes
)

type blake2b struct {
	// Represents the length of the hash
	len C.size_t
	// Represents crypto_generichash_blake2b_state, which must be 64 byte aligned.
	// This is not enforced by Go, so 64 extra bytes are allocated and
	// the 384 aligned bytes in them are used.
	state1 [384 + 64]byte
}

// SumBlake2b returns the Blake2b hash of input data `in` in output buffer `out`.
// A key `key` can be given to create a hash unique to that key.
// The length of the hash is determined by the length of the output buffer,
// which has to be between BytesMin and BytesMax (inclusive).
// The size of `key` can either be 0 or between KeyBytesMin and KeyBytesMax (inclusive).
func SumBlake2b(out, in, key []byte) {
	support.CheckSizeInRange(out, Blake2bBytesMin, Blake2bBytesMax, "out")

	if len(key) > 0 {
		support.CheckSizeInRange(key, Blake2bKeyBytesMin, Blake2bKeyBytesMax, "key")
	}

	C.crypto_generichash_blake2b(
		(*C.uchar)(support.BytePointer(out)), C.size_t(len(out)),
		(*C.uchar)(support.BytePointer(in)), C.ulonglong(len(in)),
		(*C.uchar)(support.BytePointer(key)), C.size_t(len(key)))
}

// SumBlake2bSaltPersonal returns the Blake2b salted and personalised hash of input data `in` in output buffer `out`.
// A key `key` can be given to create a hash unique to that key.
// The length of the hash is determined by the length of the output buffer,
// which has to be between BytesMin and BytesMax (inclusive).
// The size of `key` can either be 0 or between KeyBytesMin and KeyBytesMax (inclusive).
func SumBlake2bSaltPersonal(out, in, key, salt, personal []byte) {
	support.CheckSizeInRange(out, Blake2bBytesMin, Blake2bBytesMax, "out")
	support.CheckSize(salt, Blake2bSaltBytes, "salt")
	support.CheckSize(personal, Blake2bPersonalBytes, "personal")

	if len(key) > 0 {
		support.CheckSizeInRange(key, Blake2bKeyBytesMin, Blake2bKeyBytesMax, "key")
	}

	C.crypto_generichash_blake2b_salt_personal(
		(*C.uchar)(support.BytePointer(out)), C.size_t(len(out)),
		(*C.uchar)(support.BytePointer(in)), C.ulonglong(len(in)),
		(*C.uchar)(support.BytePointer(key)), C.size_t(len(key)),
		(*C.uchar)(&salt[0]),
		(*C.uchar)(&personal[0]))
}

// NewBlake2b returns a new hash.Sum for computing the Blake2b hash.
// A key `key` can be given to create a hash unique to that key.
// `size` determines the length of the hash and has to be between BytesMin and BytesMax.
// The size `key` can either be 0 or between KeyBytesMin and KeyBytesMax (inclusive).
func NewBlake2b(size int, key []byte) hash.Hash {
	support.CheckIntInRange(size, Blake2bBytesMin, Blake2bBytesMax, "hash size")

	if len(key) > 0 {
		support.CheckSizeInRange(key, Blake2bKeyBytesMin, Blake2bKeyBytesMax, "key")
	}

	s := &blake2b{
		len: C.size_t(size),
	}

	C.crypto_generichash_blake2b_init(s.state(),
		(*C.uchar)(support.BytePointer(key)),
		C.size_t(len(key)),
		s.len)

	return s
}

// NewBlake2bSaltPersonal returns a new hash.Sum for computing the salted and personalised Blake2b hash.
// A key `key` can be given to create a hash unique to that key.
// `size` determines the length of the hash and has to be between BytesMin and BytesMax.
// The size `key` can either be 0 or between KeyBytesMin and KeyBytesMax (inclusive).
func NewBlake2bSaltPersonal(size int, key, salt, personal []byte) hash.Hash {
	support.CheckIntInRange(size, Blake2bBytesMin, Blake2bBytesMax, "hash size")
	support.CheckSize(salt, Blake2bSaltBytes, "salt")
	support.CheckSize(personal, Blake2bPersonalBytes, "personal")

	if len(key) > 0 {
		support.CheckSizeInRange(key, Blake2bKeyBytesMin, Blake2bKeyBytesMax, "key")
	}

	s := &blake2b{
		len: C.size_t(size),
	}

	C.crypto_generichash_blake2b_init_salt_personal(s.state(),
		(*C.uchar)(support.BytePointer(key)),
		C.size_t(len(key)),
		s.len,
		(*C.uchar)(&salt[0]),
		(*C.uchar)(&personal[0]))

	return s
}

// state returns a pointer to the space allocated for the state
func (s *blake2b) state() *C.crypto_generichash_blake2b_state {
	var offset uintptr
	mod := uintptr(unsafe.Pointer(&s.state1)) % 64

	if mod == 0 {
		offset = mod
	} else {
		offset = 64 - mod
	}

	return (*C.crypto_generichash_blake2b_state)(unsafe.Pointer(&s.state1[offset]))
}

// Write adds data to the running hash.
func (s *blake2b) Write(p []byte) (int, error) {
	C.crypto_generichash_blake2b_update(s.state(),
		(*C.uchar)(support.BytePointer(p)),
		C.ulonglong(len(p)))

	return len(p), nil
}

// Sum returns the calculated hash appended to `b`.
func (s *blake2b) Sum(b []byte) []byte {
	out := append(b, make([]byte, s.len)...)

	C.crypto_generichash_blake2b_final(s.state(),
		(*C.uchar)(&out[len(b)]),
		s.len)

	return out
}

// Reset resets the hash to its initial state.
func (s *blake2b) Reset() {
	panic("This hash can not be reset")
}

// Size returns the number of bytes Sum will return.
func (s *blake2b) Size() int {
	return int(s.len)
}

// Block size returns the underlying block size.
// This is not exposed by libsodium, so it returns 1.
func (s *blake2b) BlockSize() int {
	return 1
}

// GenerateKeyBlake2b generates a key for use with Blake2b
func GenerateKeyBlake2b() []byte {
	k := make([]byte, Blake2bKeyBytes)
	C.crypto_generichash_blake2b_keygen((*C.uchar)(&k[0]))
	return k
}
