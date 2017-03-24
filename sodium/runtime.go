package sodium

// #cgo pkg-config: libsodium
// #include <stdlib.h>
// #include <sodium.h>
import "C"

// RuntimeHasNEON returns true if NEON is supported
func RuntimeHasNEON() bool {
	return C.sodium_runtime_has_neon() != 0
}

// RuntimeHasSSE2 returns true if SSE2 is supported
func RuntimeHasSSE2() bool {
	return C.sodium_runtime_has_sse2() != 0
}

// RuntimeHasSSE3 returns true if SSE3 is supported
func RuntimeHasSSE3() bool {
	return C.sodium_runtime_has_sse3() != 0
}

// RuntimeHasSSSE3 returns true if SSSE3 is supported
func RuntimeHasSSSE3() bool {
	return C.sodium_runtime_has_ssse3() != 0
}

// RuntimeHasSSE41 returns true if SSE4.1 is supported
func RuntimeHasSSE41() bool {
	return C.sodium_runtime_has_sse41() != 0
}

// RuntimeHasAVX returns true if AVX is supported
func RuntimeHasAVX() bool {
	return C.sodium_runtime_has_avx() != 0
}

// RuntimeHasAVX2 returns true if AVX2 is supported
func RuntimeHasAVX2() bool {
	return C.sodium_runtime_has_avx2() != 0
}

// RuntimeHasPCLMUL returns true if PCLMUL is supported
func RuntimeHasPCLMUL() bool {
	return C.sodium_runtime_has_pclmul() != 0
}

// RuntimeHasAESNI returns true if AES-NI is supported
func RuntimeHasAESNI() bool {
	return C.sodium_runtime_has_aesni() != 0
}
