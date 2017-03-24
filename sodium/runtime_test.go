package sodium

import "testing"

func TestRuntime(t *testing.T) {
	// Test Init
	if Init() {
		t.Log("Sodium initialized.")
	} else {
		t.Log("Sodium was already initialized.")
	}

	// Show runtime support
	t.Logf("Runtime support:\n"+
		"NEON:   %v\n"+
		"SSE2:   %v\n"+
		"SSE3:   %v\n"+
		"SSSE3:  %v\n"+
		"SSE41:  %v\n"+
		"AVX:    %v\n"+
		"AVX2:   %v\n"+
		"PCLMUL: %v\n"+
		"AESNI:  %v\n",
		RuntimeHasNEON(),
		RuntimeHasSSE2(),
		RuntimeHasSSE3(),
		RuntimeHasSSSE3(),
		RuntimeHasSSE41(),
		RuntimeHasAVX(),
		RuntimeHasAVX2(),
		RuntimeHasPCLMUL(),
		RuntimeHasAESNI(),
	)
}
