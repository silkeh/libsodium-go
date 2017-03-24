package sodium

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"

// Init initialises libsodium
func Init() bool {
	result := C.sodium_init()

	if result == -1 {
		panic("Sodium initialization failed")
	}

	return result == 0
}
